package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/mgo.v2"
	"time"
)

const adminPwHashBytes = 64

func GenerateAdminPassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, adminPwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

// NewUser alloc and initialize a user.
func NewAdmin(r *RegisterAdminForm, t int64) (a *Admin, err error) {
	hash, err := GenerateAdminPassHash(r.Password, beego.AppConfig.String("salt::admin"))
	if err != nil {
		return nil, err
	}

	admin := Admin{
		Username:  r.Username,
		Password:  hash,
		GenerateAccount: 	r.GenerateAccount,
		CreateArticle: 		r.CreateArticle,
		ReadArticle: 		r.ReadArticle,
		DeleteArticle: 		r.DeleteArticle,
		CreateFinancial: 	r.CreateFinancial,
		ReadFinancial: 		r.ReadFinancial,
		DeleteFinancial: 	r.DeleteFinancial,
		CreateSys: 			r.CreateSys,
		ReadSys: 			r.ReadSys,
		DeleteSys: 			r.DeleteSys,
		ReadRewards: 		r.ReadRewards,
		ReadFeedback: 		r.ReadFeedback,
		IssueIncome: 		r.IssueIncome,
		ReadIco: 			r.ReadIco,
		Created: int(time.Now().Unix()),}

	return &admin, nil
}

// FindByID query a document according to input id.
func (a *Admin) FindByUsername(username string) (ok bool) {
	o := orm.NewOrm()
	exist := o.QueryTable("admin").Filter("username", username).Exist()
	return exist
}

// Insert insert a document to collection.
func (a *Admin) Insert() (code int, err error) {
	o := orm.NewOrm()
	o.Using("admin")
	id, err := o.Insert(a)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}
	return
}

func (a *Admin) ClearPass() {
	a.Password = ""
}

func (a *Admin) CheckAdminPass(adminKey string, pass string) (code int) {
	o := orm.NewOrm()
	admin := Admin{Username: adminKey}
	if o.Read(&admin, "Username") == nil {
		hash, err := GenerateAdminPassHash(pass, beego.AppConfig.String("salt::admin"))
		if err != nil {
			return ErrSystem
		}
		if admin.Password != hash {
			return ErrNotMatch
		}
		return Success
	} else {
		return ErrDatabase
	}
}

func (a *Admin) ChangePass(userKey string, newPass string) (code int, err error) {
	o := orm.NewOrm()
	newHash, err := GenerateAdminPassHash(newPass, beego.AppConfig.String("salt::admin"))
	if err != nil {
		return ErrSystem, err
	}
	num, err := o.QueryTable("admin").Filter("username", userKey).Update(orm.Params{
		"password": newHash,
	})
	if err != nil {
		beego.Debug(num)
		return ErrDatabase, nil
	}
	return Success, nil
}

func FindAdminByPage(c *Condition) (admin []*Admin, total int64, code int) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("admin").Count()
	num, err := o.QueryTable("admin").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").All(&admin, "Id",
		"Username","GenerateAccount","OperateAccount","CreateArticle","ReadArticle","DeleteArticle","CreateFinancial","ReadFinancial","DeleteFinancial",
		"CreateSys","ReadSys","DeleteSys","ReadRewards","ReadFeedback","IssueIncome","ReadIco","Created")
	if err != nil {
		beego.Debug(num)
		return nil, 0, ErrDatabase
	}
	return admin, cnt, Success
}

func CheckPermission(username string, key string) (ok bool, code int) {
	o := orm.NewOrm()
	admin := Admin{Username: username}

	if o.Read(&admin, "Username") == nil {
		if key == "GenerateAccount" {
			return admin.GenerateAccount, Success
		} else if key == "OperateAccount" {
			return admin.OperateAccount, Success
		} else if key == "CreateArticle" {
			return admin.CreateArticle, Success
		} else if key == "ReadArticle" {
			return admin.ReadArticle, Success
		} else if key == "DeleteArticle" {
			return admin.DeleteArticle, Success
		} else if key == "CreateFinancial" {
			return admin.CreateFinancial, Success
		} else if key == "ReadFinancial" {
			return admin.ReadFinancial, Success
		} else if key == "DeleteFinancial" {
			return admin.DeleteFinancial, Success
		} else if key == "CreateSys" {
			return admin.CreateSys, Success
		} else if key == "ReadSys" {
			return admin.ReadSys, Success
		} else if key == "DeleteSys" {
			return admin.DeleteSys, Success
		} else if key == "ReadRewards" {
			return admin.ReadRewards, Success
		} else if key == "ReadInvest" {
			return admin.ReadRewards, Success
		} else if key == "ReadFeedback" {
			return admin.ReadFeedback, Success
		} else if key == "IssueIncome" {
			return admin.IssueIncome, Success
		} else if key == "ReadIco" {
			return admin.IssueIncome, Success
		} else {
			return false, ErrInput
		}
	} else {
		return false, ErrDatabase
	}
}
