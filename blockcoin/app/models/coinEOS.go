package models

import (

	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type AccountEOS struct {
	Id   		   int    		  `orm:"column(id)"`
	AccountName    string      	  `orm:"column(account_name)"`
	PrivateKey     string      	  `orm:"column(private_key)"`
	PublicKey      string      	  `orm:"column(public_key)"`
	CreatedAt      time.Time      `orm:"auto_now_add;type(datetime)"`
	UpdatedAt      time.Time      `orm:"auto_now;type(datetime)"`
}


func (u *AccountEOS) TableName() string {
	return "account_eos"
}


//把eos账户写入数据库
func InsertDatabaseEOSAccount(data CreateEOSAccount) {

	var o = orm.NewOrm()     //定义查询器
	o.Using("default")
	account_eos := new(AccountEOS)
	account_eos.AccountName = data.NewAccount
	account_eos.PublicKey   = data.NewAccountPubKey
	account_eos.PrivateKey  = data.NewAccountPriKey

	res, err := o.Insert(account_eos)

	fmt.Println(res)
	fmt.Println(err)
}