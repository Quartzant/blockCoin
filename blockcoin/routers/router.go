package routers

import (
	"blockcoin/app/controllers"
	"blockcoin/app/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	//coin
	beego.Router("/xrp/flow", &api.XRPController{})
	beego.Router("/usdt/flow", &api.USDTController{})
	beego.Router("/btc/flow", &api.BTCController{})
	beego.Router("/eth/flow", &api.ETHController{})
	beego.Router("/eos/flow", &api.EOSController{})

    //user,get发送email，同时写入session
	beego.Router("/email/sendCode", &api.UserController{},"get:SendCode")
	//注册监听
	beego.Router("/listen/register", &api.UserController{}, "post:RegisterListen")
	beego.Router("/listen/updated", &api.UserController{}, "post:UpdateUser")
	//关闭监听
	beego.Router("/listen/close", &api.UserController{}, "get:CloseListen")


	//打开监听
	beego.Router("/btc/cron", &api.BTCController{}, "get:OpenListen")
	beego.Router("/eos/cron", &api.EOSController{}, "get:OpenListen")
	beego.Router("/xrp/cron", &api.XRPController{}, "get:OpenListen")
	beego.Router("/eth/cron", &api.ETHController{}, "get:OpenListen")
	beego.Router("/usdt/cron", &api.USDTController{}, "get:OpenListen")

	//eos创建账户，交易，账户查重， 获取公私钥，购买ram, 出售ram， 获得账户信息， 查看余额， 抵押cpu和net, 查看计算net和cpu的价格
	//投票与撤销
	beego.Router("/eos/createAccount", &api.EOSController{},"post:CreateEOSAccountFunc")
	beego.Router("/eos/transfer", &api.EOSController{},"post:EOSTransfer")
	beego.Router("/eos/check", &api.EOSController{},"get:EOSAccountCheck")
	beego.Router("/eos/getKeys", &api.EOSController{},"get:EOSGetKeys")
	beego.Router("/eos/buyRam", &api.EOSController{},"post:BuyRam")
	beego.Router("/eos/sellRam", &api.EOSController{},"post:SellRam")
	//beego.Router("/eos/getAccount", &api.EOSController{},"get:EOSAccount")
	beego.Router("/eos/getAccountBalance", &api.EOSController{},"get:EOSAccountBalance")
	beego.Router("/eos/delegateBW", &api.EOSController{},"post:DelegateBW")
	beego.Router("/eos/unDelegateBW", &api.EOSController{},"post:UnDelegateBW")
	beego.Router("/eos/addPermission", &api.EOSController{},"post:AddPermissions")
	beego.Router("/eos/deletePermission", &api.EOSController{},"post:DeletePermissions")
	beego.Router("/eos/vote", &api.EOSController{},"post:Vote")
	beego.Router("/eos/unVote", &api.EOSController{},"post:UnVote")
    //发行代币
	beego.Router("/eos/createEosio", &api.EOSController{},"post:CreateEosio")

	//btc获得公私钥（暂不写入数据库），交易。
	beego.Router("/btc/getKeys", &api.BTCController{},"get:CreateBTCAccountFunc")
	beego.Router("/btc/transfer", &api.BTCController{},"post:BTCTransfer")
	//usdt
	beego.Router("/usdt/transfer", &api.USDTController{},"post:USDTTransfer")

	//eth创建地址，交易，获得余额。
	beego.Router("/wallet/getKeys", &controllers.WalletController{},"get:Create")
	beego.Router("/eth/transfer", &controllers.WalletController{},"post:TransferEth")
	//beego.Router("/eth/getBalance", &controllers.WalletController{},"get:GetEthBalance")


}
