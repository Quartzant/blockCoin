package api

import (
	"blockcoin/app/models"
	"blockcoin/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)


type UserController struct {
	beego.Controller
}

type UpdatedListenData struct {
	email    string
	code     string
	address  models.RegisterAddress
}



//注册用户信息
func (c *UserController) RegisterListen() {
	//新建用户，写入表单信息信息
	f := models.RegisterData{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &f)


	idefications := models.InsertIdentification(f.Form.Address)

	var o = orm.NewOrm()
	var listen models.CronModel

	listen.Email = f.Form.Email
	listen.Phone = f.Form.Phone
	listen.AddressBTC = f.Form.Address.AddressBTC
	listen.AddressETH = f.Form.Address.AddressETH
	listen.AddressXRP = f.Form.Address.AddressXRP
	listen.AccountEOS = f.Form.Address.AccountEOS
	listen.BtcIdentification  = idefications.BTCIdentification
	listen.EthIdentification  = idefications.ETHIdentification
	listen.XrpIdentification  = idefications.XRPIdentification
	listen.UsdtIdentification = idefications.USDTIdentification
	listen.EosIdentification  = idefications.EOSIdentification
	num, _ := o.Insert(&listen)    		         				//这里需要传递指针
	if num == 0 {
		c.Data["json"] = common.ErrInsert
		c.ServeJSON()
		return
	}

	c.Data["json"] = f
	c.ServeJSON()
	return
}

//获得邮箱验证码
func (c *UserController) SendCode() {
	//根据email写入session加上随机的code
	email := c.GetString("email")

	var emailData common.EmailData

	code := common.GetRandomNum(6)

	//写入session
	sess, _ := common.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)

	mySession := sess.Get(email)

	if mySession == nil {
		sess.Set(email, code)
		emailData.ReceiverAddress = email
		emailData.NickName  =  email
		emailData.Subject   = "验证码信息"
		emailData.Body      = "您的验证码是： " + code
		common.SendEMail(emailData)

	} else {
		sess.Delete(email)
		sess.Set(email, code)
	}
	c.Data["json"] = sess.Get(email)
	c.ServeJSON()
	return
}

//更新用户信息
func (c *UserController) UpdateUser() {
	//获得要更改的信息
	a := UpdatedListenData{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &a)


	if a.email == "" {
		c.Data["json"] = common.MissEmail
		c.ServeJSON()
		return
	}

	//根据email找session，判断code是否相同
	sess, _ := common.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	sessCode := sess.Get(a.email)
	if sessCode == nil {
		c.Data["json"] = common.MissCode
		c.ServeJSON()
		return
	} else {
		//如果是的话，允许更改
		addrIdentifacitons := models.InsertIdentification(a.address)

		// 写入数据库update
		o := orm.NewOrm()
		num, err := o.QueryTable("cron_job").Update(orm.Params{
			"btc_identification" : addrIdentifacitons.BTCIdentification,
			"eth_identification" : addrIdentifacitons.ETHIdentification,
			"xrp_identification" : addrIdentifacitons.XRPIdentification,
			"usdt_identification": addrIdentifacitons.USDTIdentification,
			"eos_identification" : addrIdentifacitons.EOSIdentification,
		})

		if num == 0 {
			c.Data["json"] = err
			c.ServeJSON()
			return
		}

		c.Data["json"] = "写入成功"
		c.ServeJSON()
		return
	}

}

//关闭监听
func (c *UserController) CloseListen(){
	email := c.GetString("email")
	switch_coin := c.GetString("switch")  //switch_btc, switch_eos...

	if email == "" {
		c.Data["json"] = common.ErrCloseCron
		c.ServeJSON()
		return
	}
	var o = orm.NewOrm()     //定义查询器
	//var user cron.CronModel		//定义容器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		switch_coin: "1",
	})

	if num != 1 {
		c.Data["json"] = common.ErrCloseCron2
		c.ServeJSON()
		return
	}
	c.Data["json"] = common.SuccessCloseCron
	c.ServeJSON()
	return
}

//执行btc的监听
func (c *BTCController) OpenListen() {
	email := c.GetString("email")

	var o = orm.NewOrm()     //定义查询器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		"switch_btc": "0",
	})
	if num != 1 {
		c.Data["json"] = common.ErrOpenCron
		c.ServeJSON()
		return
	}

	for {
		//判断是否关闭状态，是的话break
		judge := models.IfCloseListen(email, 1)
		if judge == "1" {
			break
		}
		time.Sleep(models.BTCTerminal)


		//判断是否有变
		var o = orm.NewOrm()    //定义查询器
		var user models.CronModel //定义容器
		qs := o.QueryTable("cron_job").Filter("email", email)
		qs.One(&user)   //把查出来的东西放进容器

		//获取数据
		address := user.AddressBTC

		if address == "" {
			c.Data["json"] = common.MissAddress
			c.ServeJSON()
			return
		}
		page := "1"
		pagesize := "10"
		res, err := models.GetBTC(address, page, pagesize)

		if err != nil {
			c.Data["json"] = common.ErrCron
			c.ServeJSON()
			return
		}

		//如果有变
		if 	res.TotalCount != user.BtcIdentification {

			num := res.TotalCount - user.BtcIdentification

			//发送消息
			models.SendBTCMessage(res, num, user.Email)

			//修改成不该发送状态
			qs.Update(orm.Params{
				"btc_identification": res.TotalCount,
			})
		}
		fmt.Println("没有变化")
	}

	fmt.Println("监听结束")
	c.Data["json"] = map[string]interface{}{"result": "监听结束"}
	c.ServeJSON()
	return
}

//执行eos的监听
func (c *EOSController) OpenListen() {
	email := c.GetString("email")

	var o = orm.NewOrm()     //定义查询器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		"switch_eos": "0",
	})

	if num != 1 {
		c.Data["json"] = common.ErrOpenCron
		c.ServeJSON()
		return
	}

	for {
		//判断是否关闭状态，是的话break
		judge := models.IfCloseListen(email, 2)
		if judge == "1" {
			break
		}
		time.Sleep(models.EOSTerminal)

		//判断是否有变
		var o = orm.NewOrm()    //定义查询器
		var user models.CronModel //定义容器
		qs := o.QueryTable("cron_job").Filter("email", email)
		qs.One(&user)   //把查出来的东西放进容器

		//获取数据
		address := user.AccountEOS
		page := "1"
		pagesize := "10"

		res, err := models.GetEOS(address, page, pagesize)

		if err != nil {
			c.Data["json"] = common.ErrCron
			c.ServeJSON()
			return
		}

		//如果有变
		if 	res.TraceCount != user.EosIdentification {

			num := res.TraceCount - user.EosIdentification

			//发送消息
			models.SendEOSMessage(res, num, user.Email)

			//修改成不该发送状态
			qs.Update(orm.Params{
				"eos_identification": res.TraceCount,
			})
		} else {
			fmt.Println("没有变化")
		}
	}

	fmt.Println("监听结束")
	c.Data["json"] = map[string]interface{}{"result": "监听结束"}
	c.ServeJSON()
	return
}

//执行xrp的监听
func (c *XRPController) OpenListen(){
	email := c.GetString("email")

	var o = orm.NewOrm()     //定义查询器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		"switch_xrp": "0",
	})

	if num != 1 {
		c.Data["json"] = common.ErrOpenCron
		c.ServeJSON()
		return
	}

	for {
		//判断是否关闭状态，是的话break
		judge := models.IfCloseListen(email, 5)
		if judge == "1" {
			break
		}
		time.Sleep(models.XRPTerminal)

		//判断是否有变
		var o = orm.NewOrm()    //定义查询器
		var user models.CronModel //定义容器
		qs := o.QueryTable("cron_job").Filter("email", email)
		qs.One(&user)   //把查出来的东西放进容器

		//获取数据
		address := user.AddressXRP
		count := "10"

		res, err := models.GetXRP(address, count)

		if err != nil {
			c.Data["json"] = common.ErrCron
			c.ServeJSON()
			return
		}

		//如果有变
		if 	res[0].Hash != user.XrpIdentification {

			//发送消息
			models.SendXRPMessage(res, user.XrpIdentification, user.Email)

			//修改成不该发送状态
			qs.Update(orm.Params{
				"xrp_identification": res[0].Hash,
			})
		} else {
			fmt.Println("没有变化")
		}
	}

	fmt.Println("监听结束")
	c.Data["json"] = map[string]interface{}{"result": "监听结束"}
	c.ServeJSON()
	return
}

//执行eth的监听
func (c *ETHController) OpenListen(){
	email := c.GetString("email")

	var o = orm.NewOrm()     //定义查询器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		"switch_eth": "0",
	})

	if num != 1 {
		c.Data["json"] = common.ErrOpenCron
		c.ServeJSON()
		return
	}

	for {
		//判断是否关闭状态，是的话break
		judge := models.IfCloseListen(email, 4)
		if judge == "1" {
			break
		}
		time.Sleep(models.ETHTerminal)

		//判断是否有变
		var o = orm.NewOrm()    //定义查询器
		var user models.CronModel //定义容器
		qs := o.QueryTable("cron_job").Filter("email", email)
		qs.One(&user)   //把查出来的东西放进容器

		//获取数据
		address := user.AddressETH
		limit := "10"

		res, err := models.GetETH(address, limit)

		if err != nil {
			c.Data["json"] = common.ErrCron
			c.ServeJSON()
			return
		}

		//如果有变
		if 	res[0].Hash != user.EthIdentification {

			//发送消息
			models.SendETHMessage(res, user.EthIdentification, user.Email)

			//修改成不该发送状态
			qs.Update(orm.Params{
				"eth_identification": res[0].Hash,
			})
		} else {
			fmt.Println("没有变化")
		}
	}

	fmt.Println("监听结束")
	c.Data["json"] = map[string]interface{}{"result": "监听结束"}
	c.ServeJSON()
	return
}

//执行usdt的监听
func (c *USDTController) OpenListen(){
	email := c.GetString("email")

	var o = orm.NewOrm()     //定义查询器
	num, _ := o.QueryTable("cron_job").Filter("email", email).Update(orm.Params{
		"switch_usdt": "0",
	})

	if num != 1 {
		c.Data["json"] = common.ErrOpenCron
		c.ServeJSON()
		return
	}

	for {
		//判断是否关闭状态，是的话break
		judge := models.IfCloseListen(email, 3)
		if judge == "1" {
			break
		}
		time.Sleep(models.USDTTerminal)

		//判断是否有变
		var o = orm.NewOrm()    //定义查询器
		var user models.CronModel //定义容器
		qs := o.QueryTable("cron_job").Filter("email", email)
		qs.One(&user)   //把查出来的东西放进容器

		//获取数据
		address := user.AddressBTC

		res, err := models.GetUSDT(address)

		if err != nil {
			c.Data["json"] = common.ErrCron
			c.ServeJSON()
			return
		}

		//如果有变
		if 	res[0].Txid != user.UsdtIdentification {

			//发送消息
			models.SendUSDTMessage(res, user.UsdtIdentification, user.Email)

			//修改成不该发送状态
			qs.Update(orm.Params{
				"usdt_identification": res[0].Txid,
			})
		} else {
			fmt.Println("没有变化")
		}
	}

	fmt.Println("监听结束")
	c.Data["json"] = map[string]interface{}{"result": "监听结束"}
	c.ServeJSON()
	return
}






