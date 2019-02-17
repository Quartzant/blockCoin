package eoscli

import (
	"fmt"

	"github.com/eoscanada/eos-go"
)

func GetCurrencyBalance(api *eos.API, account eos.AccountName, symbol string, tokenAccount eos.AccountName) {
	balances, err := api.GetCurrencyBalance(account, symbol, tokenAccount)
	if err == nil {
		for _, asset := range balances {
			fmt.Printf("%s\n", asset)
		}
	}
}


func GetCurrencyBalanceReturn(api *eos.API, account eos.AccountName, symbol string, tokenAccount eos.AccountName) ([]eos.Asset, error) {
	balances, err := api.GetCurrencyBalance(account, symbol, tokenAccount)
	if err == nil {
		for _, asset := range balances {
			fmt.Printf("%s\n", asset)
		}
	}

	return balances, err
}