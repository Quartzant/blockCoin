package eoscli

import (
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
)

func GetAccount(api *eos.API, account eos.AccountName) {
	resp, _ := api.GetAccount(account)
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}

//func GetAccountReturn(api *eos.API, account eos.AccountName) (string, error) {
//	resp, err := api.GetAccount(account)
//	beego.Debug(err)
//	fmt.Println(resp)
//	data, _ := json.MarshalIndent(resp, "", "  ")
//	fmt.Println(string(data))
//	return string(data), err
//}
