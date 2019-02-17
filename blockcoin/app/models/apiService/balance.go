package apiService

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"

	"io/ioutil"
)

func (as *apiService) BtcBalance(address string) (balance interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	btcBalance, err := redis.String(r.Do("GET", "balance:btc:"+address))

	if err != nil {
		res, err := as.request("GET", "balance?active=" + address, params,false, false, false)
		if err != nil {
			return balance, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return balance, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return balance, err
		}
		if err := json.Unmarshal(textRes, &balance); err != nil {
			return balance, err
		}
		saveBtcbalance, err := redis.String(r.Do("SET", "balance:btc:"+address, textRes, "EX", "5"))
		if len(saveBtcbalance) == 0 {
			return balance, err
		}
		if err != nil {
			return balance, err
		}
		defer r.Close()
		return balance, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(btcBalance), &balance); err != nil {
			return balance, err
		}
		return balance, nil
	}
}

func (as *apiService) EthBalance(address string) (balance interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	ethBalance, err := redis.String(r.Do("GET", "balance:eth:"+address))

	if err != nil {
		res, err := as.request("GET", "api?module=account&action=balance&address=" + address + "&tag=latest&apikey=", params,false, false, false)
		if err != nil {
			return balance, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return balance, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return balance, err
		}
		if err := json.Unmarshal(textRes, &balance); err != nil {
			return balance, err
		}
		saveEthbalance, err := redis.String(r.Do("SET", "balance:eth:"+address, textRes, "EX", "5"))
		if len(saveEthbalance) == 0 {
			return balance, err
		}
		if err != nil {
			return balance, err
		}
		defer r.Close()
		beego.Debug(1)
		return balance, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(ethBalance), &balance); err != nil {
			return balance, err
		}
		beego.Debug(2)
		return balance, nil
	}
}


