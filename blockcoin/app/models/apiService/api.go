package apiService

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Service interface {
	Ticker24() ([]*QuotesResultOutput, error)
	Letter20() ([]*LettersResultOutput, error)
	InquireAddressTx(tx TxRequest) (result []string, err error)
	InquireTxDetails(tx TxDetailsRequest) (result TxDetailsOutput, code int)
	InquireIcoTxDetails(tx IcoTxDetailsRequest) (result IcoTxDetailsOutput, code int)
	BtcFee() (fee interface{}, err error)
	EthFee() (fee interface{}, err error)
	BtcBalance(address string) (balance interface{}, err error)
	EthBalance(address string) (balance interface{}, err error)
	EthTx(address string) (balance interface{}, err error)
	BtcTx(address string) (balance interface{}, err error)
	Live() (live interface{}, err error)
	Quotes() (quotes interface{}, err error)
	BtcPrice() (price Price, err error)
	EthPrice() (price Price, err error)
}

// Service represents service layer for Binance API.
//
// The main purpose for this layer is to be replaced with dummy implementation
// if necessary without need to replace Binance instance.

// NewAPIService creates instance of Service.
//
// If logger or ctx are not provided, NopLogger and Background context are used as default.
// You can use context for one-time request cancel (e.g. when shutting down the app).
func NewAPIService(url string, ctx context.Context) Service {
	if ctx == nil {
		ctx = context.Background()
	}
	return &apiService{
		URL:    url,
		Ctx:    ctx,
	}
}

func (as *apiService) request(method string, endpoint string, params map[string]string,
	apiKey bool, sign bool, p bool) (*http.Response, error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	transport := &http.Transport{}
	if p {
		transport = &http.Transport{Proxy: proxy}
	}

	client := &http.Client{Transport: transport}

	url := fmt.Sprintf("%s/%s", as.URL, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	req.WithContext(as.Ctx)

	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}