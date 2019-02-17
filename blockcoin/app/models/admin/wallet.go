package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/mgo.v2"
	"time"
)

func (f *FinancialMan) Insert() (int) {
	o := orm.NewOrm()
	o.Using("financial_man")
	id, err := o.Insert(f)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func (a *Feedback) Insert() (int) {
	o := orm.NewOrm()
	o.Using("feedback")
	id, err := o.Insert(a)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		beego.Debug(err)
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func (s *SystemNotification) Insert() (int) {
	o := orm.NewOrm()
	o.Using("system_notification")
	id, err := o.Insert(s)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func (i *InvestmentList) Insert() (int) {
	o := orm.NewOrm()
	o.Using("investment_list")
	id, err := o.Insert(i)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func (i *Ico) Insert() (int) {
	o := orm.NewOrm()
	o.Using("ico")
	id, err := o.Insert(i)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func (a *Article) Insert() (int) {
	o := orm.NewOrm()
	o.Using("article")
	id, err := o.Insert(a)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	if err != nil {
		if mgo.IsDup(err) {
			return ErrDupRows
		} else {
			return ErrDatabase
		}
	} else {
		return Success
	}
}

func FindFinancialMan(c *Condition) ([]*FinancialMan, int64, int){
	o := orm.NewOrm()
	var output []*FinancialMan
	cnt, err := o.QueryTable("financial_man").Count()
	num, err := o.QueryTable("financial_man").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").All(&output, "Id", "Token", "Term", "Min", "IncomeRate", "ReceiptAddress")
	if err != nil {
		beego.Debug(num)
		return nil, 0, ErrDatabase
	}
	return output, cnt, Success
}

func FindAllFinancialMan() ([]*FinancialMan, int){
	o := orm.NewOrm()
	var output []*FinancialMan
	num, err := o.QueryTable("financial_man").OrderBy("id").All(&output, "Id", "Token", "Term", "Min", "IncomeRate", "ReceiptAddress")
	if err != nil {
		beego.Debug(num)
		return nil, ErrDatabase
	}
	return output, Success
}

func IsMatchTokenAndAddress(f *FinancialMan) (code int) {
	if f.Token == "eth" {
		if len(f.ReceiptAddress) != 42 {
			return ErrNotMatch
		} else {
			return Success
		}
	} else if f.Token == "btc" {
		if len(f.ReceiptAddress) == 34 || len(f.ReceiptAddress) == 35 {
			return Success
		} else {
			return ErrNotMatch
		}
	} else {
		return ErrInput
	}
}

func IsExistFinancialMan(f *FinancialMan) (ok bool) {
	o := orm.NewOrm()
	exist:= o.QueryTable("financial_man").Filter("token", f.Token).Filter("term", f.Term).Filter("min", f.Min).Filter("income_rate", f.IncomeRate).Filter("receipt_address", f.ReceiptAddress).Exist()
	return exist
}

func DeleteFinancialManById(id int) (code int) {
	o := orm.NewOrm()
	beego.Debug(id)
	if num, err := o.Delete(&FinancialMan{Id: id}); err != nil {
		return ErrDatabase
	} else if num == 0 {
		return ErrInput
	} else {
		return Success
	}
}

func FindAllFeedback(c *Condition) (feedback []*Feedback, total int64, code int){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("feedback").Count()
	num, err := o.QueryTable("feedback").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").All(&feedback)
	if err != nil {
		beego.Debug(num)
		return nil, 0, ErrDatabase
	}
	return feedback, cnt, Success
}

func IsExistSysNotification(s *SystemNotification) (ok bool) {
	o := orm.NewOrm()
	exist:= o.QueryTable("system_notification").Filter("title", s.Title).Filter("content", s.Content).Exist()
	return exist
}

func FindAllSystemNotification(c *Condition) (sysNotification []*SystemNotification, total int64, code int){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("system_notification").Count()
	num, err := o.QueryTable("system_notification").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").All(&sysNotification)
	if err != nil {
		beego.Debug(num)
		return nil, 0, ErrDatabase
	}
	return sysNotification, cnt, Success
}

func DeleteSysNotificationById(id int) (code int) {
	o := orm.NewOrm()
	if num, err := o.Delete(&SystemNotification{Id: id}); err != nil {
		return ErrDatabase
	} else {
		fmt.Println(num)
		return Success
	}
}

func FindAddressMinReceipt(address string) (float32, error){
	o := orm.NewOrm()
	f := FinancialMan{ReceiptAddress: address}

	if err := o.Read(&f, "ReceiptAddress"); err == nil {
		return f.Min, nil
	} else {
		return 0, err
	}
}

func FindFinancialById(id int) (*FinancialMan, error) {
	o := orm.NewOrm()
	f := FinancialMan{Id: id}

	if err := o.Read(&f); err == nil {
		return &f, nil
	} else {
		return nil, err
	}
}

func FindInvestmentByAddress(address string, c *Condition) (i []*InvestmentList, result []string, num int64, code int) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("investment_list").Filter("ReceiptAddress", address).Count()
	num, err = o.QueryTable("investment_list").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").Filter("ReceiptAddress", address).All(&i)
	if err != nil {
		beego.Debug(num)
		return nil, nil, 0, ErrDatabase
	}
	for _, item := range i {
		result = append(result, item.Tx)
	}
	return i, result, cnt, Success
}

func QueryAllMaturityRewards(address string, c *Condition) (i []*InvestmentList, num int64, code int){
	o := orm.NewOrm()

	cnt, err := o.QueryTable("investment_list").Filter("ReceiptAddress", address).Count()
	num, err = o.QueryTable("investment_list").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").Filter("end_time__lte", time.Now().Unix()).Filter("ReceiptAddress", address).All(&i)
	if err != nil {
		beego.Debug(num)
		return nil, 0, ErrDatabase
	}
	return i, cnt, Success
}

func UpdateRewards(id int) (code int) {
	o := orm.NewOrm()

	num, err := o.QueryTable("investment_list").Filter("id", id).Update(orm.Params{
		"state": 1,
	})
	fmt.Println(num)
	if err != nil {
		return ErrDatabase
	}
	return Success
}

func FindAllArticles(c *Condition) (article []*Article, total int64, code int){
	o := orm.NewOrm()
	cnt, err := o.QueryTable("article").Count()
	num, err := o.QueryTable("article").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("-id").All(&article)
	if err != nil {
		return nil, 0, ErrDatabase
	}
	beego.Debug(num)
	return article, cnt, Success
}

func DeleteArticleById(id int) (code int) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Article{Id: id}); err != nil {
		return ErrDatabase
	} else {
		fmt.Println(num)
		return Success
	}
}

func DeleteAdminById(id int) (code int) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Admin{Id: id}); err != nil {
		return ErrDatabase
	} else {
		fmt.Println(num)
		return Success
	}
}

func FindIcoDetails(c *Condition) (ico []*Ico, result []string, total int64, code int) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("ico").Count()
	num, err := o.QueryTable("ico").Limit(c.Limit).Offset((c.Offset -1)*c.Limit).OrderBy("id").All(&ico, "Id", "Tx", "Address", "Sum", "Created")
	if err != nil {
		beego.Debug(num)
		return nil, result, 0, ErrDatabase
	}
	for _, item := range ico {
		result = append(result, item.Tx)
	}
	return ico, result, cnt, Success
}