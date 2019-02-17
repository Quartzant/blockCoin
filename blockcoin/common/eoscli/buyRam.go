package eoscli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

func BuyRam(api *eos.API, payer, receiver eos.AccountName, numBytes int) {
	action := system.NewBuyRAMBytes(payer, receiver, uint32(numBytes))
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	fmt.Println(err)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}


func BuyRamReturn(api *eos.API, payer, receiver eos.AccountName, numBytes int) (string, error) {
	action := system.NewBuyRAMBytes(payer, receiver, uint32(numBytes))
	actions := []*eos.Action{action}
	resp, err := api.SignPushActions(actions...)
	data, _ := json.MarshalIndent(resp, "", "  ")

	return string(data), err
}