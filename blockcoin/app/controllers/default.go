package controllers

import (
	"blockcoin/app/models/jwtUserModel"
	"blockcoin/common"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	token := c.GetString("token")
	appid, err := common.Token_auth(token, "secret")
	if err != nil {

		c.Data["json"] = common.ErrExpired
		c.ServeJSON()
		return
	}
	roleid, err := jwtUserModel.Auth_role("2321fd", appid)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = common.ErrPermission
		c.ServeJSON()
		return
	}
	fmt.Println(roleid)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
