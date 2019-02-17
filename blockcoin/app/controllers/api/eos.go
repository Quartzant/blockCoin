package api

import (
	"blockcoin/app/models"
	"blockcoin/common"
	"blockcoin/common/eoscli"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
	"reflect"
	"regexp"
	"strings"
)



//eos创建账户
func (c *EOSController) CreateEOSAccountFunc() {
	//创建后返回公钥私钥，保存在数据库
	data := models.CreateEOSAccount{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	creator := eos.AN(data.CreatorAccount)
	models.EOSInit(data.CreatorPriKey)
	new_account := eos.AN(data.NewAccount)

	a, errPub := ecc.NewPublicKey(data.NewAccountPubKey)	//验证公钥
	_, errPri := ecc.NewPrivateKey(data.NewAccountPriKey)	//验证私钥

	b, _ := regexp.MatchString("^[0-5a-z]{12,12}$", data.NewAccount)


	if b == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}

	if errPub != nil {
		c.Data["json"] = common.ErrPubKey
		c.ServeJSON()
		return
	}
	if errPri != nil {
		c.Data["json"] = common.ErrPriKey
		c.ServeJSON()
		return
	}
	fmt.Println(new_account)
	fmt.Println(creator)
	fmt.Println(data)

	ram_quantity, _ := eos.NewEOSAssetFromString(data.Quantity.RamQuantity)
	cpu_quantity, _ := eos.NewEOSAssetFromString(data.Quantity.CpuQuantity)
	net_quantity, _ := eos.NewEOSAssetFromString(data.Quantity.NetQuantity)

	fmt.Println(ram_quantity)

	fmt.Println(cpu_quantity)

	fmt.Println(net_quantity)

	response, err := eoscli.NewAccountReturn(models.EosApi, creator,
		new_account, a, ram_quantity, cpu_quantity, net_quantity, true)
	beego.Debug(err)

	if response != "" {
		res := make(map[string]interface{})
		res["newAccount"] = data.NewAccount
		c.Data["json"] = res

		//写入数据库
		//models.InsertDatabaseEOSAccount(data)

	} else {
		res := make(map[string]interface{})
		res["error"] = err
		c.Data["json"] = res
	}

	c.ServeJSON()
	return
}

//eos转账交易
func (c *EOSController) EOSTransfer() {
	data := models.TransferEOS{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	models.EOSInit(data.FromPriKey)  //初始化

	from := eos.AN(data.FromAccount)
	to := eos.AN(data.ToAccount)
	quantity, _ := eos.NewEOSAssetFromString(data.Quantity)
	memo := data.Memo
	code := eos.AN(data.Code)

	ret, err := eoscli.TransferReturn(models.EosApi, from, to, quantity, memo, code)
	res := make(map[string]interface{})

	if err != nil {
		res["error"] = err
		res["message"] = "转账失败，请根据error信息查询原因"
	} else {
		res["data"] = ret
		res["message"] = "转账成功"
	}
	c.Data["json"] = res
	c.ServeJSON()
	return
}

//检测账户是否存在
func (c *EOSController) EOSAccountCheck() {
	var account string
	res := make(map[string]interface{})
	account = c.GetString("account_eos")
	api := eos.New(beego.AppConfig.String("eos_rpc"))

	b, _ := regexp.MatchString("^[0-5a-z]{12,12}$", account)


	if b == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}
	eoscli.GetAccount(api, eos.AN(account))
	accountResp, err := api.GetAccount(eos.AN(account))

	fmt.Println(accountResp)
	fmt.Println(err)
	if err != nil {

		res["data"] = "该账户可以注册"

	} else {
		res["data"] = "该账户已被注册"
	}
	res["message"] = accountResp
	c.Data["json"] = res
	c.ServeJSON()
	return
}

//生成符合EOS账户的公私钥
func (c *EOSController) EOSGetKeys() {
	priKey, pubKey, err :=eoscli.NewKeysReturn()

	if err != nil {
		c.Data["json"] = common.ErrGetKeys
		c.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	data["priKey"] = priKey
	data["pubKey"] = pubKey
	c.Data["json"] = data
	c.ServeJSON()
	return
}

//获得余额balance
func (c *EOSController) EOSAccountBalance() {
	models.EOSInit("")
	account := c.GetString("account_eos")
	symbol := c.GetString("symbol")
	fmt.Println(account)
	fmt.Println(symbol)
	if account == ""  {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}

	data, _ := eoscli.GetCurrencyBalanceReturn(models.EosApi, eos.AN(account), symbol, eos.AN("eosio.token"))

	if len(data) == 0{
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}
	fmt.Println(reflect.TypeOf(data))
	c.Data["json"] = data
	c.ServeJSON()
	return
}

//购买ram
func (c *EOSController) BuyRam() {
	var buyRamData models.EOSBuyRam
	json.Unmarshal(c.Ctx.Input.RequestBody, &buyRamData)

	if err := c.VerifyForm(&buyRamData); err != nil {
		beego.Debug(err)
		c.Data["json"] = common.MissParameter
		c.ServeJSON()
		return
	}

	a, _ := regexp.MatchString("^[0-5a-z]{12,12}$", buyRamData.ToAccount)
	b, _ := regexp.MatchString("^[0-5a-z]{12,12}$", buyRamData.FromAccount)


	if b == false || a == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}

	models.EOSInit(buyRamData.PrivateKey)

	data ,err := eoscli.BuyRamReturn(models.EosApi, eos.AN(buyRamData.FromAccount),
		eos.AN(buyRamData.ToAccount), buyRamData.NumBytes)

	if err != nil {
		c.Data["json"] = common.ErrBuyRam
		c.ServeJSON()
		return
	}


	var res models.DelegateResult

	err2 := json.Unmarshal([]byte(data), &res)
	beego.Debug(err2)

	c.Data["json"] = res
	c.ServeJSON()
	return
}

//出售ram
func (c *EOSController) SellRam() {
	var sellRamData models.EOSSellRam

	json.Unmarshal(c.Ctx.Input.RequestBody, &sellRamData)

	if err := c.VerifyForm(&sellRamData); err != nil {
		beego.Debug(err)
		c.Data["json"] = common.MissParameter
		c.ServeJSON()
		return
	}

	a, _ := regexp.MatchString("^[0-5a-z]{12,12}$", sellRamData.Account)


	if  a == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}

	models.EOSInit(sellRamData.PrivateKey)

	data, err := eoscli.SellRamReturn(models.EosApi, eos.AN(sellRamData.Account), sellRamData.NumBytes)
	beego.Debug(err)
	if err != nil {
		c.Data["json"] = common.ErrSellRam
		c.ServeJSON()
		return
	}

	var res models.DelegateResult

	err2 := json.Unmarshal([]byte(data), &res)
	beego.Debug(err2)

	c.Data["json"] = res
	c.ServeJSON()
	return
}

//抵押net，cpu
func (c *EOSController) DelegateBW() {
	var delegateData models.EOSDelegateBW

	json.Unmarshal(c.Ctx.Input.RequestBody, &delegateData)

	if err := c.VerifyForm(&delegateData); err != nil {
		beego.Debug(err)
		c.Data["json"] = common.MissParameter
		c.ServeJSON()
		return
	}

	a, _ := regexp.MatchString("^[0-5a-z]{12,12}$", delegateData.ToAccount)
	b, _ := regexp.MatchString("^[0-5a-z]{12,12}$", delegateData.FromAccount)


	if b == false || a == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}


	models.EOSInit(delegateData.PrivateKey)
	cpuStake, _ := eos.NewAsset(delegateData.SkateCpuQuantity)
	netStake, _ := eos.NewAsset(delegateData.SkateNetQuantity)
	data, err := eoscli.DelegateBWReturn(models.EosApi, eos.AN(delegateData.FromAccount),
		eos.AN(delegateData.ToAccount), cpuStake, netStake, delegateData.Transfer)

	if err != nil {
		c.Data["json"] = common.ErrDelegateBW
		c.ServeJSON()
		return
	}

	var res models.DelegateResult

	err2 := json.Unmarshal([]byte(data), &res)
	beego.Debug(err2)

	c.Data["json"] = res
	c.ServeJSON()
	return
}

//赎回net，cpu
func (c *EOSController) UnDelegateBW() {
	var unDelegateData models.EOSUnDelegateBW

	json.Unmarshal(c.Ctx.Input.RequestBody, &unDelegateData)

	if err := c.VerifyForm(&unDelegateData); err != nil {
		beego.Debug(err)
		c.Data["json"] = common.MissParameter
		c.ServeJSON()
		return
	}

	a, _ := regexp.MatchString("^[0-5a-z]{12,12}$", unDelegateData.ToAccount)
	b, _ := regexp.MatchString("^[0-5a-z]{12,12}$", unDelegateData.FromAccount)


	if b == false || a == false {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}


	models.EOSInit(unDelegateData.PrivateKey)
	cpuStake, _ := eos.NewAsset(unDelegateData.SkateCpuQuantity)
	netStake, _ := eos.NewAsset(unDelegateData.SkateNetQuantity)
	data, err := eoscli.UndelegateBWReturn(models.EosApi, eos.AN(unDelegateData.FromAccount),
		eos.AN(unDelegateData.ToAccount), cpuStake, netStake)

	if err != nil {
		c.Data["json"] = common.ErrDelegateBW
		c.ServeJSON()
		return
	}

	var res models.DelegateResult

	err2 := json.Unmarshal([]byte(data), &res)
	beego.Debug(err2)

	c.Data["json"] = res
	c.ServeJSON()
	return
}

//添加permissions
func (c *EOSController) AddPermissions() {
	var addPermissionData models.AddPermission

	json.Unmarshal(c.Ctx.Input.RequestBody, &addPermissionData)


	re, _ := regexp.MatchString("^[0-5a-z]{12,12}$", addPermissionData.Account)
	if re == false  {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}


	models.EOSInit(addPermissionData.PrivateKey)

	a := eos.Authority{}

	a.Threshold = uint32(addPermissionData.Threshold)

	pubkey, _ := ecc.NewPublicKey(addPermissionData.NewPublicKey)
	a.Keys = []eos.KeyWeight{}

	var middle eos.KeyWeight
	middle.PublicKey = pubkey
	middle.Weight = uint16(addPermissionData.Weight)

	a.Keys = append(a.Keys, middle)


	b := system.NewUpdateAuth(eos.AN(addPermissionData.Account), eos.PN(addPermissionData.NewPermissionName),
		eos.PN(addPermissionData.ParentPermissionName), a, eos.PN(addPermissionData.ParentPermissionName))

	_, err1:=models.EosApi.SignPushActions(b)
	if err1 != nil {
		c.Data["json"] = common.ErrAddPermission
		c.ServeJSON()
		return
	}

	beego.Debug(err1)
	fmt.Println(err1)
	c.Data["json"] = b.Data
	c.ServeJSON()
	return
}

//删除
func (c *EOSController) DeletePermissions() {
	var deletePermission models.DeletePermission


	json.Unmarshal(c.Ctx.Input.RequestBody, &deletePermission)


	re, _ := regexp.MatchString("^[0-5a-z]{12,12}$", deletePermission.Account)
	if re == false  {
		c.Data["json"] = common.ErrAccount
		c.ServeJSON()
		return
	}

	models.EOSInit(deletePermission.PrivateKey)

	b := system.NewDeleteAuth(eos.AN(deletePermission.Account), eos.PN(deletePermission.PermissionName))

	_, err1 :=models.EosApi.SignPushActions(b)


	if err1 != nil {
		c.Data["json"] = common.ErrDelPermission
		c.ServeJSON()
		return
	}

	beego.Debug(err1)
	fmt.Println(err1)

	c.Data["json"] = b.Data
	c.ServeJSON()
	return
}





func (c *EOSController) Vote() {

	a := strings.ToTitle("i am the blond of my sword")

	fmt.Println(a)

}

func (c *EOSController) UnVote() {

	//
}


func (this *EOSController) CreateEosio() {
	//部署代币合约


	//创建代币
	models.EOSInit("5JWVSgFSx94mce6FYbdP29Y5ouct776eDmgsEQRQ1Lam56d8AEt")
	issure := eos.AN("shinechain12")

	var EOSSymbol = eos.Symbol{Precision: 5, Symbol: "JOEGE"}
	s := eos.Asset{100.0000, EOSSymbol}

	action := token.NewCreate(issure, s)


	_, err1:=models.EosApi.SignPushActions(action)


	if err1 != nil {
		beego.Debug(err1)
		this.Data["json"] = err1
		this.ServeJSON()
		return
	}

	beego.Debug(err1)
	fmt.Println(err1)


	this.Data["json"] = action.Data
	this.ServeJSON()
	return
}