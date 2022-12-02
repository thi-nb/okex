package public_data

import (
	"github.com/thi-nb/okex/okex/models/publicdata"
)

type (
	GetInstruments struct {
		Code        int                      `json:"code,string"`
		Msg         string                   `json:"msg,omitempty"`
		Instruments []*publicdata.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		Code      int                                   `json:"code,string"`
		Msg       string                                `json:"msg,omitempty"`
		Histories []*publicdata.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		Code          int                        `json:"code,string"`
		Msg           string                     `json:"msg,omitempty"`
		OpenInterests []*publicdata.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		Code         int                       `json:"code,string"`
		Msg          string                    `json:"msg,omitempty"`
		FundingRates []*publicdata.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		Code        int                      `json:"code,string"`
		Msg         string                   `json:"msg,omitempty"`
		LimitPrices []*publicdata.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		Code             int                            `json:"code,string"`
		Msg              string                         `json:"msg,omitempty"`
		OptionMarketData []*publicdata.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Code                            int                                          `json:"code,string"`
		Msg                             string                                       `json:"msg,omitempty"`
		EstimatedDeliveryExercisePrices []*publicdata.EstimatedDeliveryExercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Code                                 int                                               `json:"code,string"`
		Msg                                  string                                            `json:"msg,omitempty"`
		GetDiscountRateAndInterestFreeQuotas []*publicdata.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		Code        int                      `json:"code,string"`
		Msg         string                   `json:"msg,omitempty"`
		SystemTimes []*publicdata.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		Code              int                            `json:"code,string"`
		Msg               string                         `json:"msg,omitempty"`
		LiquidationOrders []*publicdata.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		Code       int                     `json:"code,string"`
		Msg        string                  `json:"msg,omitempty"`
		MarkPrices []*publicdata.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		Code          int                        `json:"code,string"`
		Msg           string                     `json:"msg,omitempty"`
		PositionTiers []*publicdata.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		Code                      int                                    `json:"code,string"`
		Msg                       string                                 `json:"msg,omitempty"`
		InterestRateAndLoanQuotas []*publicdata.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		Code       int        `json:"code,string"`
		Msg        string     `json:"msg,omitempty"`
		Underlings [][]string `json:"data,omitempty"`
	}
	Status struct {
		Code   int                `json:"code,string"`
		Msg    string             `json:"msg,omitempty"`
		States []publicdata.State `json:"data,omitempty"`
	}
)
