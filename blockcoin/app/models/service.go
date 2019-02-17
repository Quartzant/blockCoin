package models

import (
	"blockcoin/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/eoscanada/eos-go"
	"strconv"
	"strings"
	"time"
)

var (
	EosApi *eos.API
)

const (
	SwitchBTC   = 1
	SwitchEOS   = 2
	SwitchUSDT  = 3
	SwitchETH   = 4
	SwitchXRP   = 5
)


//判断关闭监听
func IfCloseListen(email string, coinType int) string {
	//如果数据库某个值为0者break
	var s = orm.NewOrm()
	var sw CronModel
	s.QueryTable("cron_job").Filter("email", email).One(&sw) //把查出来的东西放进容器

	if sw.SwitchBTC  == 1 && coinType == SwitchBTC {
		return "1"
	}
	if sw.SwitchEOS  == 1 && coinType == SwitchEOS {
		return "1"
	}
	if sw.SwitchUSDT == 1 && coinType == SwitchUSDT {
		return "1"
	}
	if sw.SwitchETH  == 1 && coinType == SwitchETH {
		return "1"
	}
	if sw.SwitchXRP  == 1 && coinType == SwitchXRP {
		return "1"
	}
	return "0"
}

//定义发送消息内容
func SendBTCMessage(list BtcData, num int, email string) {
	fmt.Println("您的BTC地址有", num, "条新的转账消息")
	count := num

	var newData string
	if num >= 10 {
		num = 10
	}
	for i := 0; i< num; i++{
		fmt.Println(list.List[i])
		newData = newData + "第"+ strconv.Itoa(i+1) +"条的哈希是"+list.List[i].Hash
	}
	var message = common.EmailData{}
	message.ReceiverAddress = email
	message.NickName        = "新的BTC转账消息"
	message.Subject         = "您的BTC地址有" + strconv.Itoa(count) + "条新的转账记录"
	if count > 10 {
		message.Body            = newData + "(最多显示最新的10条记录，更多请查看交易流水)"
	} else {
		message.Body            = newData
	}

	common.SendEMail(message)
}


func SendEOSMessage(list EosData, num int, email string) {
	fmt.Println("您的EOS账户有", num, "条新的转账消息")
	count := num

	var newData string
	if num > 10 {
		num = 10
	}

	for i := 0; i< num; i++{
		fmt.Println(list.List[i])
		newData = newData + "第"+ strconv.Itoa(i+1) +"条的哈希是"+list.List[i].TrxId
	}
	var message = common.EmailData{}
	message.ReceiverAddress = email
	message.NickName        = "新的EOS转账消息"
	message.Subject         = "您的EOS地址有" + strconv.Itoa(count) + "条新的转账记录"
	if count > 10 {
		message.Body            = newData + "(最多显示最新的10条记录，更多请查看交易流水)"
	} else {
		message.Body            = newData
	}
	common.SendEMail(message)
}


func SendXRPMessage(list []XrpTransactions, beforehash string, email string) {

	count := 0
	var newData string

	for i := 0; i< len(list); i++{

		if list[i].Hash == beforehash {
			break
		} else {
			fmt.Println(list[i])
			newData = newData + "第"+ strconv.Itoa(i+1) +"条的哈希是"+list[i].Hash
		}
		count = count+1
	}

	fmt.Println("您的XRP地址有", count, "条新的转账消息")
	var message = common.EmailData{}
	message.ReceiverAddress = email
	message.NickName        = "新的XRP转账消息"
	if count == 10 {
		message.Subject         = "您的XRP地址有" + strconv.Itoa(count) + "条以上新的转账记录"
		message.Body            = newData + "(最多显示最新的10条，更多请查看交易流水)"
	} else {
		message.Subject         = "您的XRP地址有" + strconv.Itoa(count) + "条新的转账记录"
		message.Body            = newData
	}
	common.SendEMail(message)
}


func SendETHMessage(list []EthCoin, beforehash string, email string) {

	count := 0
	var newData string

	for i := 0; i< len(list); i++{

		if list[i].Hash == beforehash {
			break
		} else {
			fmt.Println(list[i])
			newData = newData + "第"+ strconv.Itoa(i+1) +"条的哈希是"+list[i].Hash
		}
		count = count+1
	}

	fmt.Println("您的ETH地址有", count-1, "条新的转账消息")
	var message = common.EmailData{}
	message.ReceiverAddress = email
	message.NickName        = "新的ETH转账消息"
	if count == 10 {
		message.Subject         = "您的ETH地址有" + strconv.Itoa(count) + "条以上的新转账记录"
		message.Body            = newData + "(最多显示最新的10条，更多请查看交易流水)"
	} else {
		message.Subject         = "您的ETH地址有" + strconv.Itoa(count) + "条新的转账记录"
		message.Body            = newData
	}


	common.SendEMail(message)
}


func SendUSDTMessage(list []UsdtTransactions, beforehash string, email string) {

	count := 0
	var newData string

	for i := 0; i< len(list); i++{

		if list[i].Txid == beforehash {
			break
		} else {
			fmt.Println(list[i])
			newData = newData + "第"+ strconv.Itoa(i+1) +"条的哈希是"+list[i].Txid
		}
		count = count+1
	}

	fmt.Println("您的USDT有", count, "条新的转账消息")
	var message = common.EmailData{}
	message.ReceiverAddress = email
	message.NickName        = "新的USDT转账消息"
	message.Subject         = "您的USDT地址有" + strconv.Itoa(count) + "条新的转账记录"
	message.Body            = newData
	common.SendEMail(message)
}


//eos初始化私钥，倒入的私钥用来决定这次动作的行为主体是哪个账户
func EOSInit(priKey string){
	rpc := beego.AppConfig.String("eos_rpc")
	key := priKey
	EosApi = eos.New(rpc)
	signer := eos.NewKeyBag()
	signer.ImportPrivateKey(key)
	EosApi.SetSigner(signer)
}

//接收BTC的unspent信息
func BTCUnspent(address string)(BTCUnspentOutputs, error){
	//测试网
	//url := "https://testnet.blockchain.info/unspent?active="+address

	//正式网
	url := "https://blockchain.info/unspent?active="+address
	req := httplib.Get(url)

	str, err := req.String()

	var btcunspent BTCUnspentOutputs
	json.Unmarshal([]byte(str), &btcunspent)

	return btcunspent, err
}

//BTC和USDT签名
func Sign(tx *wire.MsgTx, privKeyStr string, prevPkScripts [][]byte)  {
	inputs := tx.TxIn
	wif, err := btcutil.DecodeWIF(privKeyStr)

	fmt.Println("wif err", err)
	privKey := wif.PrivKey

	for i := range inputs {
		pkScript := prevPkScripts[i]
		var script []byte
		script, err = txscript.SignatureScript(tx, i, pkScript, txscript.SigHashAll,
			privKey, true)
		inputs[i].SignatureScript = script
	}
}

//构造utxo, 构造outputs
func GetUnspents (address string, amount int64, outputs []*wire.TxOut) []*wire.TxOut {

	addr, err1 := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err1 != nil {
		common.WriteToLog("decodeAddress", err1)
	}
	pkScript, err2 := txscript.PayToAddrScript(addr)


	if err2 != nil {
		common.WriteToLog("makePayScript", err2)
	}

	outputs =  append(outputs, wire.NewTxOut(amount, pkScript))

	return outputs
}

//BTC和USDT交易hex广播到公网
func PushBTCTX (hex string) (string, error) {

	url := "https://blockchain.info/pushtx?tx=" + hex

	req := httplib.Post(url)

	str, err := req.String()

	return str, err
}

//获得五种币的数据流水
func GetBTC(address string, page string, pagesize string) (BtcData, interface{}) {
	//address := beego.AppConfig.String("btc_address")

	url := "https://chain.api.btc.com/v3/address/"+address+"/tx?page="+ page + "&pagesize=" + pagesize
	req := httplib.Get(url)
	str, err2 := req.String()
	fmt.Println(req)
	fmt.Println(str)
	fmt.Println(err2)
	var btc BtcCoin

	err := json.Unmarshal([]byte(str), &btc)

	for i := 0; i < len(btc.Data.List); i++ {
		btc.Data.List[i].Fee = btc.Data.List[i].Fee * 0.000000001
		btc.Data.List[i].BalanceDiff = btc.Data.List[i].BalanceDiff * 0.000000001
	}


	return btc.Data, err
}


func GetETH(address string, limit string) ([]EthCoin, interface{}) {
	//address := beego.AppConfig.String("eth_address")

	url := "http://api.ethplorer.io/getAddressTransactions/"+address+"?apiKey=freekey&limit=" + limit
	req := httplib.Get(url)
	str, _ := req.String()

	var eth []EthCoin

	err := json.Unmarshal([]byte(str), &eth)

	return	eth, err
}


func GetEOS(account string, page string, size string) (EosData, interface{}) {
	//account := beego.AppConfig.String("eos_account")

	url := "https://api.eospark.com/api?module=account&action=get_account_related_trx_info&" +
		"apikey=a9564ebc3289b7a14551baf8ad5ec60a&" +
		"account="+ account +
		"&page=" + page +
		"&size=" + size +
		"&symbol=EOS" +
		"&code=eosio.token "
	req := httplib.Get(url)

	str, _ := req.String()

	var eosVal EosCoin

	err := json.Unmarshal([]byte(str), &eosVal)

	for i := 0; i < len(eosVal.Data.List); i++ {
		eosVal.Data.List[i].Timestamp = TransEosTime(eosVal.Data.List[i].Timestamp)
	}

	return	eosVal.Data, err
}


func GetUSDT(address string) ([]UsdtTransactions, interface{}) {

	req := httplib.Post("https://api.omniexplorer.info/v1/transaction/address")
	//req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("addr",address)
	req.Param("page","0")



	var usdt UsdtCoin

	str, err := req.String()
	fmt.Println(err)
	err2 := json.Unmarshal([]byte(str), &usdt)
	fmt.Println(err2)
	return usdt.Transactions, err
}


func GetXRP(address string, count string) ([]XrpTransactions, interface{}) {

	url := "https://data.ripple.com/v2/accounts/"+ address +"/transactions?descending=true&type=Payment&limit=" + count


	req := httplib.Get(url)


	str, _ := req.String()
	var xrp XrpCoin
	err := json.Unmarshal([]byte(str), &xrp)

	for i := 0; i < len(xrp.Transactions); i++ {
		v1, err1 := strconv.ParseFloat(xrp.Transactions[i].Tx.Fee, 64)
		v2, err2 := strconv.ParseFloat(xrp.Transactions[i].Tx.Amount, 64)

		if err1 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
		}

		xrp.Transactions[i].Tx.Fee= strconv.FormatFloat(v1 * 0.000001, 'f', -1, 32)
		xrp.Transactions[i].Tx.Amount= strconv.FormatFloat(v2* 0.000001, 'f', -1, 32)
	}
	fmt.Println(err)
	return xrp.Transactions, err
}


//功能函数
func TransEosTime(timeAccept string) string {
	str01 := strings.Replace(timeAccept, "T", " ", 1)
	str02 := strings.Replace(str01, ".500", "", 1)
	str03 := strings.Replace(str02, ".000", "", 1)
	base_format := "2006-01-02 15:04:05"
	t, _ := time.Parse(base_format, str03)
	ant := t.Add(8*3600*1e9)  //+8小时
	str_time := ant.Format(base_format)
	return str_time
}
