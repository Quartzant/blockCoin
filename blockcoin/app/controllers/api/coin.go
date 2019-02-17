package api

import (
	"blockcoin/app/models"
	"blockcoin/common"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"strings"
)




type XRPController struct {
	common.BaseController
}

type USDTController struct {
	common.BaseController
}

type BTCController struct {
	common.BaseController
}

type ETHController struct {
	common.BaseController
}

type EOSController struct {
	common.BaseController
}



//获得历史流水
func (c *XRPController) Get() {
	address := c.GetString("address")
	count    := c.GetString("count")

	if address == "" {
		address = beego.AppConfig.String("xrp_address")
	}

	if count == "" {
		count = "20"
	}


	res, err := models.GetXRP(address, count)

	if err != nil {
		c.Data["json"] = common.ErrGetMessage
		c.ServeJSON()
		return
	}

	c.Data["json"] = res
	c.ServeJSON()
	return
}

func (c *USDTController) Get() {
	address := c.GetString("address")

	if address =="" {

		address = beego.AppConfig.String("btc_address")
	}

	res, err := models.GetUSDT(address)
	fmt.Println(err)
	if err != nil {
		c.Data["json"] = common.ErrGetMessage
		c.ServeJSON()
		return
	}

	c.Data["json"] = res
	c.ServeJSON()
	return
}

func (c *BTCController) Get() {
	address  := c.GetString("address")
	page     := c.GetString("page")
	pagesize := c.GetString("pagesize")

	if address =="" {
		address = beego.AppConfig.String("btc_address")
	}
	if page =="" {
		page = "1"
	}
	if pagesize =="" {
		pagesize = "20"
	}
	fmt.Println(address, page, pagesize)
	res, err := models.GetBTC(address, page, pagesize)
	fmt.Println(res, err)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = common.ErrGetMessage
		c.ServeJSON()
		return
	}

	c.Data["json"] = res.List
	c.ServeJSON()
	return
}

func (c *ETHController) Get() {
	address  := c.GetString("address")
	limit    := c.GetString("limit")

	if address =="" {
		address = beego.AppConfig.String("eth_address")
	}
	if limit =="" {
		limit = "20"
	}

	res, err := models.GetETH(address, limit)

	if err != nil {
		c.Data["json"] = common.ErrGetMessage
		c.ServeJSON()
		return
	}

	c.Data["json"] = res
	c.ServeJSON()
	return
}

func (c *EOSController) Get() {
	account  := c.GetString("account")
	page     := c.GetString("page")
	size 	 := c.GetString("size")

	if account =="" {
		account = beego.AppConfig.String("eos_account")
	}
	if page =="" {
		page = "1"
	}
	if size =="" {
		size = "20"
	}

	res, err := models.GetEOS(account, page, size)

	if err != nil {
		c.Data["json"] = common.ErrGetMessage
		c.ServeJSON()
		return
	}

	c.Data["json"] = res.List
	c.ServeJSON()
	return
}







//生成BTC地址
func (c *BTCController) CreateBTCAccountFunc() {

	wifKey, address, _ := models.GenerateBTCTest() // 测试地址
	//wifKey2, address2, _ := models.GenerateBTC() // 正式地址

	a := make(map[string]interface{})

	a["wifKey"] = wifKey
	a["address"] = address
	//a["wifKey2"] = wifKey2
	//a["address2"] = address2
	c.Data["json"] = a
	c.ServeJSON()
	return
}

//BTC交易
func (c *BTCController) BTCTransfer() {

	var f models.BTCTransferReceive
	json.Unmarshal(c.Ctx.Input.RequestBody, &f)
	unspentData, err := models.BTCUnspent(f.FromAddress)      //获得unspent的数据


	//传f，unspentData

	if err != nil {
		c.Data["json"] = common.NotEnough
		c.ServeJSON()
		return
	}

	//计算Satoshi
	cost := f.Amount + f.Fee              //总花费

	var totalUnspent int64  			  //总unspent
	for _, item := range unspentData.UnspentOutputs {
		totalUnspent = totalUnspent + item.Value
	}

	leftToMe := totalUnspent - cost       //留给自己的

	//钱不够
	if leftToMe < 0 {
		c.Data["json"] = common.NotEnough
		c.ServeJSON()
		return
	}
	///////

	//构造输出
	outputs := []*wire.TxOut{}
	if leftToMe != 0 {
		outputs = models.GetUnspents(f.FromAddress, leftToMe, outputs)      //给自己
	}
	outputs = models.GetUnspents(f.ToAddress, f.Amount, outputs)		     //给别人


	//构造输入 （没有排序）(循环扣除unspent)
	prevPkScripts := make([][]byte, len(unspentData.UnspentOutputs))
	inputs := make([]*wire.TxIn, 0)

	for i := 0; i < len(unspentData.UnspentOutputs); i++ {
		prevTxHash := unspentData.UnspentOutputs[i].TxHashBigEndian
		prevPkScriptHex := unspentData.UnspentOutputs[i].Script
		prevTxOutputN := uint32(unspentData.UnspentOutputs[i].TxOutputN)
		fmt.Println(prevTxHash)
		fmt.Println(prevPkScriptHex)
		fmt.Println(prevTxOutputN)

		hash, _ := chainhash.NewHashFromStr(prevTxHash) // tx hash
		outPoint := wire.NewOutPoint(hash, prevTxOutputN) // 第几个输出

		txIn := wire.NewTxIn(outPoint, nil, nil)
		inputs = append(inputs, txIn)

		prevPkScript, _ := hex.DecodeString(prevPkScriptHex)
		prevPkScripts[i] = prevPkScript
	}


	tx := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}

	//签名
	privKey := f.Private
	models.Sign(tx, privKey, prevPkScripts)



	//获得hex，发送上公网
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
	}
	txHex := hex.EncodeToString(buf.Bytes())
	fmt.Println("hextest", txHex)

	res, err := models.PushBTCTX(txHex)
	result := make(map[string]interface{})

	if err != nil {
		result["message"] = common.ErrInSystem
	} else {
		result["message"] = res
	}
	c.Data["json"] = txHex
	c.ServeJSON()
	return
}

//usdt交易
func (c *USDTController) USDTTransfer() {

	var f models.USDTTransferReceive
	json.Unmarshal(c.Ctx.Input.RequestBody, &f)
	unspentData, err := models.BTCUnspent(f.FromAddress)      //获得unspent的数据

	if err != nil {
		c.Data["json"] = common.NotEnough
		c.ServeJSON()
		return
	}


	//计算Satoshi
	cost := 546 + f.Fee              //固定546sat给别人

	var totalUnspent int64  			  //总unspent
	for _, item := range unspentData.UnspentOutputs {
		totalUnspent = totalUnspent + item.Value
	}

	leftToMe := totalUnspent - cost       //留给自己的

	//钱不够
	if leftToMe < 0 {
		c.Data["json"] = common.NotEnough
		c.ServeJSON()
		return
	}

	//构造输出
	outputs := []*wire.TxOut{}
	if leftToMe != 0 {
		outputs = models.GetUnspents(f.FromAddress, leftToMe, outputs)      //给自己
	}
	outputs = models.GetUnspents(f.ToAddress, 546, outputs)		     //给别人

	//加一个omni(带着toAmount的usdt输出)
	var  toAmount int
	toAmount = int(f.Amount * 1e8)
	amountHex := common.DecimalToAny(toAmount, 16)
	usdtSend := "6f6d6e69" + "0000" + "00000000001f" + strings.Repeat("0", 16 - len(amountHex)) + amountHex
	b, _ := hex.DecodeString(usdtSend)
	fmt.Println("b=", b)
	pkScript,_ := txscript.NullDataScript(b)
	outputs = append(outputs,wire.NewTxOut(int64(0),pkScript))


	//构造输入 （没有排序）(循环扣除unspent)
	prevPkScripts := make([][]byte, len(unspentData.UnspentOutputs))
	inputs := make([]*wire.TxIn, 0)

	for i := 0; i < len(unspentData.UnspentOutputs); i++ {

		prevTxHash := unspentData.UnspentOutputs[i].TxHashBigEndian
		prevPkScriptHex := unspentData.UnspentOutputs[i].Script
		prevTxOutputN := uint32(unspentData.UnspentOutputs[i].TxOutputN)
		fmt.Println(prevTxHash)
		fmt.Println(prevPkScriptHex)
		fmt.Println(prevTxOutputN)

		hash, _ := chainhash.NewHashFromStr(prevTxHash) // tx hash
		outPoint := wire.NewOutPoint(hash, prevTxOutputN) // 第几个输出

		txIn := wire.NewTxIn(outPoint, nil, nil)
		inputs = append(inputs, txIn)

		prevPkScript, _ := hex.DecodeString(prevPkScriptHex)
		prevPkScripts[i] = prevPkScript
	}


	tx := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}

	//签名
	privKey := f.Private
	models.Sign(tx, privKey, prevPkScripts)

	//获得hex，发送上公网
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
	}
	txHex := hex.EncodeToString(buf.Bytes())
	fmt.Println("hextest", txHex)

	res, err := models.PushBTCTX(txHex)
	result := make(map[string]interface{})

	if err != nil {
		result["message"] = common.ErrInSystem
	} else {
		result["message"] = res
	}
	c.Data["json"] = txHex
	c.ServeJSON()
	return
}