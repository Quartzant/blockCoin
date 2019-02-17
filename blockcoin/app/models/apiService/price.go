package apiService

import (
	"encoding/json"
	"io/ioutil"

	"github.com/garyburd/redigo/redis"
)

type Price struct {
	Ticker		  string
	ExchangeName  string
	Base 		string
	Currency 	string
	Symbol		string
	High		float64
	Open		float64
	Close 		float64
	Low 		float64
	Vol 		float64
	Degree  	float64
	Value 		float64
	ChangeValue  	float64
	CommissionRatio float64
	QuantityRatio 	float64
	TurnoverRate 	float64
	DateTime		float64
}

func (as *apiService) BtcPrice() (price Price, err error) {
	params := make(map[string]string)
	r := RedisConn()
	btcPrice, err := redis.String(r.Do("GET", "price:btc"))

	if err != nil {
		res, err := as.request("GET", "api/v1/tick/BITFINEX:BTCUSD", params,false, false, false)
		if err != nil {
			return price, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return price, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return price, err
		}
		if err := json.Unmarshal(textRes, &price); err != nil {
			return price, err
		}
		saveBtcPrice, err := redis.String(r.Do("SET", "price:btc", textRes, "EX", "60"))
		if len(saveBtcPrice) == 0 {
			return price, err
		}
		if err != nil {
			return price, err
		}
		defer r.Close()
		return price, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(btcPrice), &price); err != nil {
			return price, err
		}
		return price, nil
	}
}

func (as *apiService) EthPrice() (price Price, err error) {
	params := make(map[string]string)
	r := RedisConn()
	ethPrice, err := redis.String(r.Do("GET", "price:eth"))

	if err != nil {
		res, err := as.request("GET", "api/v1/tick/BITFINEX:ETHUSD", params,false, false, false)
		if err != nil {
			return price, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return price, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return price, err
		}
		if err := json.Unmarshal(textRes, &price); err != nil {
			return price, err
		}
		saveEthPrice, err := redis.String(r.Do("SET", "price:eth", textRes, "EX", "60"))
		if len(saveEthPrice) == 0 {
			return price, err
		}
		if err != nil {
			return price, err
		}
		defer r.Close()
		return price, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(ethPrice), &price); err != nil {
			return price, err
		}
		return price, nil
	}
}