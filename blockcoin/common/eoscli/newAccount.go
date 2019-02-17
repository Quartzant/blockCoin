package eoscli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
)

func NewAccount(api *eos.API, creator, newAccount eos.AccountName, pubKey ecc.PublicKey, buyRAMAmount, cpuStake, netStake eos.Asset, doTransfer bool) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewNewAccount(creator, newAccount, pubKey))
	actions = append(actions, system.NewDelegateBW(creator, newAccount, cpuStake, netStake, doTransfer))
	actions = append(actions, system.NewBuyRAM(creator, newAccount, uint64(buyRAMAmount.Amount)))
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}


func NewAccountReturn(api *eos.API, creator, newAccount eos.AccountName,
	pubKey ecc.PublicKey, buyRAMAmount, cpuStake, netStake eos.Asset, doTransfer bool)(string, error) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewNewAccount(creator, newAccount, pubKey))
	actions = append(actions, system.NewDelegateBW(creator, newAccount, cpuStake, netStake, doTransfer))
	actions = append(actions, system.NewBuyRAM(creator, newAccount, uint64(buyRAMAmount.Amount)))
	resp, err1 := api.SignPushActions(actions...)

	if err1 == nil {
		data, err2 := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))

		return string(data), err2
	} else {
		return "", err1
	}
}
