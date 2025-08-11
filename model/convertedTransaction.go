package model

type ConvertedTransaction struct {
	Transaction
	ExchangeRate   float64 `json:"exchange_rate"`
	ConvertedValue float64 `json:"converted_value"`
}
