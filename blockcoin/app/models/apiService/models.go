package apiService

import (
	"context"
)

const (
	Success		= 0
	ErrDatabase = 1
	ErrSystem   = 2
	ErrWrite	= 3
	ErrDupRows  = 4
	ErrNotFound = 5
	ErrInput    = 6
	ErrNotExist = 7
	ErrNotMatch = 8
	ErrValidate = 9
	ErrCrawl	= 10
)

type apiService struct {
	URL    string
	Ctx    context.Context
}

type TxRequest struct {
	Token			string
	Address			string
}

type TxDetailsRequest struct {
	Token			string
	Tx				string
	Address			string
}

type IcoTxDetailsRequest struct {
	Tx			string
	Address		string
}

type Ticker24Body struct {
	Symbol			   string
	PriceChangePercent string
	LastPrice          string
}

type Fee struct {
	FastestFee		int
	HalfHourFee		int
	HourFee			int
}

type BtcAddress struct {

}

type TxRefsBody struct {
	Txrefs		[]struct{
		TxHash		string  `json:"tx_hash"`
	}
}

type TxDetailsBody struct {
	Confirmed	string
	Inputs		[]struct{
		OutputValue		int64  `json:"output_value"`
		Addresses		[]string
	}
	Outputs		[]struct{
		Value			int64  `json:"value"`
		Addresses		[]string
	}
}

type TxBody struct {
	TxHash 		string
}

type QuotesResultOutput struct {
	Token 				string
	PriceChangePercent  float64
	Price          		float64
}

type LettersResultOutput struct {
	Time 			string
	Title   	    string
	Content			string
	Up				string
	Down			string
}

type TxDetailsOutput struct {
	ReceiptAddress		string
	IncomeAddress		string
	Sum					float64
	Time				int64
}

type IcoTxDetailsOutput struct {
	Address				string
	Sum					float64
	Time				int64
}