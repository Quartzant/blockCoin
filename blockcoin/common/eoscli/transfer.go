package eoscli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/token"
)

//交易
func Transfer(api *eos.API, from eos.AccountName, to eos.AccountName, quantity eos.Asset, memo string, code eos.AccountName) {
	action := token.NewTransfer(from, to, quantity, memo)
	action.Account = code
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}

func TransferReturn (api *eos.API, from eos.AccountName, to eos.AccountName,
	quantity eos.Asset, memo string, code eos.AccountName) (string, error){
	action := token.NewTransfer(from, to, quantity, memo)
	action.Account = code
	actions := []*eos.Action{action}
	resp, err1 := api.SignPushActions(actions...)
	if err1 == nil {
		data, err2 := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
		return string(data), err2
	}

	fmt.Println(err1)
	return "", err1

}
