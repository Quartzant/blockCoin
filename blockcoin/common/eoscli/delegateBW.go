package eoscli

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

//用来实现抵押代币获取 cpu 和带宽资源
func DelegateBW(api *eos.API, from, receiver eos.AccountName, cpuStake, netStake eos.Asset, doTransfer bool) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewDelegateBW(from, receiver, cpuStake, netStake, doTransfer))
	resp, err := api.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}


func DelegateBWReturn(api *eos.API, from, receiver eos.AccountName,
	cpuStake, netStake eos.Asset, doTransfer bool) (string, error) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewDelegateBW(from, receiver, cpuStake, netStake, doTransfer))
	resp, err := api.SignPushActions(actions...)
	data, _ := json.MarshalIndent(resp, "", "  ")

	return string(data), err
}

