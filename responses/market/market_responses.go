package market

import (
	"github.com/thi-nb/okex/responses"
)

type (
	Ticker struct {
		Code    int           `json:"code,string"`
		Msg     string        `json:"msg,omitempty"`
		Tickers []interface{} `json:"data,omitempty"`
	}
	IndexTicker struct {
		Code         int           `json:"code,string"`
		Msg          string        `json:"msg,omitempty"`
		IndexTickers []interface{} `json:"data,omitempty"`
	}
	OrderBook struct {
		Code       int           `json:"code,string"`
		Msg        string        `json:"msg,omitempty"`
		OrderBooks []interface{} `json:"data,omitempty"`
	}
	Candle struct {
		Code    int           `json:"code,string"`
		Msg     string        `json:"msg,omitempty"`
		Candles []interface{} `json:"data,omitempty"`
	}
	IndexCandle struct {
		Code    int           `json:"code,string"`
		Msg     string        `json:"msg,omitempty"`
		Candles []interface{} `json:"data,omitempty"`
	}
	CandleMarket struct {
		Code    int           `json:"code,string"`
		Msg     string        `json:"msg,omitempty"`
		Candles []interface{} `json:"data,omitempty"`
	}
	Trade struct {
		responses.Basic
		Trades []interface{} `json:"data,omitempty"`
	}
	TotalVolume24H struct {
		responses.Basic
		TotalVolume24Hs []interface{} `json:"data,omitempty"`
	}
	IndexComponent struct {
		responses.Basic
		IndexComponents interface{} `json:"data,omitempty"`
	}
)
