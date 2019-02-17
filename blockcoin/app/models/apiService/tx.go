package apiService

import (
	"blockcoin/app/models/admin"
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"strings"
	"time"
)

func (as *apiService) InquireAddressTx(tx TxRequest) (result []string, err error) {
	params := make(map[string]string)

	res, err := as.request("GET", "/v1/" + tx.Token + "/main/addrs/" + tx.Address, params,false, false, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "a")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, err
	}

	rawTxRefs := TxRefsBody{}
	if err := json.Unmarshal(textRes, &rawTxRefs); err != nil {
		return nil, errors.Wrap(err, "rawTicker24 unmarshal failed")
	}
	if tx.Token == "eth" {
		for _, txRefs := range rawTxRefs.Txrefs {
			result = append(result, "0x" + txRefs.TxHash)
		}
	} else {
		for _, txRefs := range rawTxRefs.Txrefs {
			result = append(result, txRefs.TxHash)
		}
	}

	return result, nil
}

func (as *apiService) InquireTxDetails(tx TxDetailsRequest) (result TxDetailsOutput, code int) {
	params := make(map[string]string)

	res, err := as.request("GET", "/v1/" + tx.Token + "/main/txs/" + tx.Tx, params,false, false, false)
	if err != nil {
		return result, ErrCrawl
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, ErrCrawl
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, ErrCrawl
	}

	rawTxDetails := TxDetailsBody{}
	if err := json.Unmarshal(textRes, &rawTxDetails); err != nil {
		return result, ErrNotMatch
	}
	if len(rawTxDetails.Inputs) != 1 {
		return result, ErrNotMatch
	}
	min, err := admin.FindAddressMinReceipt(tx.Address)
	if err != nil {
		return result, ErrDatabase
	}
	t, _ := time.Parse("2006-01-02T15:04:05Z", rawTxDetails.Confirmed)
	for _, txOutputs := range rawTxDetails.Outputs {
		var sum float64
		if tx.Token == "btc" {
			sum = float64(txOutputs.Value)/1e8
			if len(txOutputs.Addresses) > 0 && tx.Address == txOutputs.Addresses[0] && float32(sum) >= min {
				result = TxDetailsOutput{
					ReceiptAddress: txOutputs.Addresses[0],
					IncomeAddress: rawTxDetails.Inputs[0].Addresses[0],
					Sum: sum,
					Time: t.Unix(),
				}
				return result, Success
			} else {
				return result, ErrInput
			}
		} else if tx.Token == "eth" {
			sum = float64(txOutputs.Value)/1e18
			if len(txOutputs.Addresses) > 0 && strings.ToLower(tx.Address) == "0x" + txOutputs.Addresses[0] && float32(sum) >= min {
				result = TxDetailsOutput{
					ReceiptAddress: txOutputs.Addresses[0],
					IncomeAddress: rawTxDetails.Inputs[0].Addresses[0],
					Sum: sum,
					Time: t.Unix(),
				}
				return result, Success
			} else {
				return result, ErrInput
			}
		} else {
			return result, ErrInput
		}

	}
	return result, ErrInput
}

func (as *apiService) InquireIcoTxDetails(tx IcoTxDetailsRequest) (result IcoTxDetailsOutput, code int) {
	params := make(map[string]string)

	res, err := as.request("GET", "/v1/eth/main/txs/" + tx.Tx, params,false, false, false)
	if err != nil {
		return result, ErrCrawl
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, ErrCrawl
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, ErrCrawl
	}

	rawTxDetails := TxDetailsBody{}
	if err := json.Unmarshal(textRes, &rawTxDetails); err != nil {
		return result, ErrNotMatch
	}
	if len(rawTxDetails.Inputs) != 1 {
		return result, ErrNotMatch
	}
	t, _ := time.Parse("2006-01-02T15:04:05Z", rawTxDetails.Confirmed)
	for _, txOutputs := range rawTxDetails.Outputs {
		sum := float64(txOutputs.Value)/1e18
		if len(txOutputs.Addresses) > 0 && strings.ToLower(tx.Address) == "0x" + txOutputs.Addresses[0] && float32(sum) >= float32(0.0001) {
			result = IcoTxDetailsOutput{
				Address: rawTxDetails.Inputs[0].Addresses[0],
				Sum: sum,
				Time: t.Unix(),
			}
			return result, Success
		} else {
			return result, ErrInput
		}
	}
	return result, ErrInput
}

func (as *apiService) EthTx(address string) (tx interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	ethTx, err := redis.String(r.Do("GET", "tx:eth:"+address))

	if err != nil {
		res, err := as.request("GET", "api?module=account&action=txlist&address=" + address + "&startblock=0&endblock=99999999&sort=desc&apikey=", params,false, false, false)
		if err != nil {
			return tx, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return tx, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return tx, err
		}
		if err := json.Unmarshal(textRes, &tx); err != nil {
			return tx, err
		}
		saveEthTx, err := redis.String(r.Do("SET", "tx:eth:"+address, textRes, "EX", "60"))
		if len(saveEthTx) == 0 {
			return tx, err
		}
		if err != nil {
			return tx, err
		}
		defer r.Close()
		return tx, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(ethTx), &tx); err != nil {
			return tx, err
		}
		return tx, nil
	}
}

func (as *apiService) BtcTx(address string) (tx interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	btcTx, err := redis.String(r.Do("GET", "tx:btc:"+address))

	if err != nil {
		res, err := as.request("GET", "rawaddr/" + address, params,false, false, false)
		if err != nil {
			return tx, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return tx, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return tx, err
		}
		if err := json.Unmarshal(textRes, &tx); err != nil {
			return tx, err
		}
		saveBtcTx, err := redis.String(r.Do("SET", "tx:btc:"+address, textRes, "EX", "60"))
		if len(saveBtcTx) == 0 {
			return tx, err
		}
		if err != nil {
			return tx, err
		}
		defer r.Close()
		return tx, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(btcTx), &tx); err != nil {
			return tx, err
		}
		return tx, nil
	}
}