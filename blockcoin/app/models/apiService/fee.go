package apiService

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"

	"io/ioutil"
)

func (as *apiService) BtcFee() (fee interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	btcFee, err := redis.String(r.Do("GET", "fee:btc"))
	if err != nil {
		res, err := as.request("GET", "api/v1/fees/recommended", params,false, false, false)
		if err != nil {
			return fee, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fee, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return fee, err
		}
		if err := json.Unmarshal(textRes, &fee); err != nil {
			return fee, err
		}
		saveBtcFee, err := redis.String(r.Do("SET", "fee:btc", textRes, "EX", "60"))
		if len(saveBtcFee) == 0 {
			return fee, err
		}
		if err != nil {
			return fee, err
		}
		defer r.Close()
		return fee, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(btcFee), &fee); err != nil {
			return fee, err
		}
		return fee, nil
	}
}

func (as *apiService) EthFee() (fee interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	ethFee, err := redis.String(r.Do("GET", "fee:eth"))

	if err != nil {
		res, err := as.request("GET", "v1/jsonrpc/mainnet/eth_gasPrice", params,false, false, false)
		if err != nil {
			return fee, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fee, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return fee, err
		}
		if err := json.Unmarshal(textRes, &fee); err != nil {
			return fee, err
		}
		saveEthFee, err := redis.String(r.Do("SET", "fee:eth", textRes, "EX", "60"))
		if len(saveEthFee) == 0 {
			return fee, err
		}
		if err != nil {
			return fee, err
		}
		defer r.Close()
		return fee, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(ethFee), &fee); err != nil {
			return fee, err
		}
		return fee, nil
	}
}