package apiService

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"

	"io/ioutil"
)

func (as *apiService) Live() (live interface{}, err error) {
	params := make(map[string]string)

	r := RedisConn()
	coindogLive, err := redis.String(r.Do("GET", "coindog:live"))
	if err != nil {
		res, err := as.request("GET", "live/list", params,false, false, false)
		if err != nil {
			return live, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return live, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return live, err
		}
		if err := json.Unmarshal(textRes, &live); err != nil {
			return live, err
		}
		coindogLive, err := redis.String(r.Do("SET", "coindog:live", textRes, "EX", "120"))
		if len(coindogLive) == 0 {
			return live, err
		}
		if err != nil {
			return live, err
		}
		defer r.Close()
		return live, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(coindogLive), &live); err != nil {
			return live, err
		}
		return live, nil
	}
}

func (as *apiService) Quotes() (quotes interface{}, err error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol
	r := RedisConn()
	coindogQuotes, err := redis.String(r.Do("GET", "coindog:quotes"))
	if err != nil {
		res, err := as.request("GET", "api/v1/currency/ranks", params,false, false, false)
		if err != nil {
			return quotes, err
		}
		textRes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return quotes, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return quotes, err
		}
		if err := json.Unmarshal(textRes, &quotes); err != nil {
			return quotes, err
		}
		coindogQuotes, err := redis.String(r.Do("SET", "coindog:quotes", textRes, "EX", "5"))
		if len(coindogQuotes) == 0 {
			return quotes, err
		}
		if err != nil {
			return quotes, err
		}
		defer r.Close()
		return quotes, nil
	} else {
		defer r.Close()
		if err := json.Unmarshal([]byte(coindogQuotes), &quotes); err != nil {
			return quotes, err
		}
		return quotes, nil
	}
}

