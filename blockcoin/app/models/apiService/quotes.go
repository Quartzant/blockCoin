package apiService

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
)

func (as *apiService) Ticker24() ([]*QuotesResultOutput, error) {
	params := make(map[string]string)
	// params["symbol"] = tr.Symbol

	res, err := as.request("GET", "api/v1/ticker/24hr", params,false, false, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Ticker/24hr")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
	}

	var result []*QuotesResultOutput

	rawAllTicker24 := []*Ticker24Body{}
	if err := json.Unmarshal(textRes, &rawAllTicker24); err != nil {
		return nil, errors.Wrap(err, "rawTicker24 unmarshal failed")
	}

	for _, rawTicker24 := range rawAllTicker24{
		if rawTicker24.Symbol == "BTCUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "BTC",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "ETHUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "ETH",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "XRPUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "XRP",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "BCCUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "BCH",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "EOSUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "EOS",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "XLMUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "XLM",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "LTCUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "LTC",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "ADAUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "ADA",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "TRXUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "TRX",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "IOTAUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "IOTA",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "BNBUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "BNB",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "NEOUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "NEO",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "ETCUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "ETC",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "ONTUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "ONT",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
		if rawTicker24.Symbol == "ZECUSDT" {
			pcPercent, err := strconv.ParseFloat(rawTicker24.PriceChangePercent, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.PriceChangePercent")
			}
			lastPrice, err := strconv.ParseFloat(rawTicker24.LastPrice, 64)
			if err != nil {
				return nil, errors.Wrap(err, "cannot parse Ticker24.LastPrice")
			}

			result = append(result, &QuotesResultOutput{
				Token: "ZEC",
				PriceChangePercent: pcPercent,
				Price:  lastPrice,
			})
		}
	}
	return result, nil
}