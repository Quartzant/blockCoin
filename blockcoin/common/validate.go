package common

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

// VerifyForm use validation to verify input parameters.
func (base *BaseController) VerifyForm(obj interface{}) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		return err
	}
	if !ok {
		str := ""
		for _, err := range valid.Errors {
			str += err.Key + ":" + err.Message + ";"
		}
		return errors.New(str)
	}

	return nil
}