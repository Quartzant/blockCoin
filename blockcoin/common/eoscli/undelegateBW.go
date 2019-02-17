package eoscli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

//用来解除抵押，释放资源，收回代币
func UndelegateBW(api *eos.API, from, receiver eos.AccountName, cpuStake, netStake eos.Asset) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewUndelegateBW(from, receiver, cpuStake, netStake))
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}

func UndelegateBWReturn(api *eos.API, from,
	receiver eos.AccountName, cpuStake, netStake eos.Asset) (string, error) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewUndelegateBW(from, receiver, cpuStake, netStake))
	resp, err := api.SignPushActions(actions...)

	data, _ := json.MarshalIndent(resp, "", "  ")

	return string(data), err
}



