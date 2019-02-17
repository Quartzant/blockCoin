package main

import (
	_ "blockcoin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func main() {

	localhost  := beego.AppConfig.String("localhost")
	port       := beego.AppConfig.String("port")
	username   := beego.AppConfig.String("username")
	pwd        := beego.AppConfig.String("password")
	database   := beego.AppConfig.String("database")

	//先连接了本地数据库，然后在models的init里面注册
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", username+":"+pwd+"@tcp("+localhost+"" +
	":"+port+")/"+database+"?charset=utf8", 30, 30)
	orm.Debug = true
	beego.Run()

}




//_ "blockcoin/routers"
//
//"github.com/astaxie/beego"
//"github.com/astaxie/beego/orm"


//localhost  := beego.AppConfig.String("localhost")
//port       := beego.AppConfig.String("port")
//username   := beego.AppConfig.String("username")
//pwd        := beego.AppConfig.String("password")
//database   := beego.AppConfig.String("database")
//
////先连接了本地数据库，然后在models的init里面注册
//orm.RegisterDriver("mysql", orm.DRMySQL)
//orm.RegisterDataBase("default", "mysql", username+":"+pwd+"@tcp("+localhost+"" +
//":"+port+")/"+database+"?charset=utf8", 30, 30)
//orm.Debug = true
//beego.Run()