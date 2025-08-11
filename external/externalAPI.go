package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ExchangeRateFetcher interface {
	FetchValidExchangeRate(date, sixMonths, currency string) (float64, error)
}

func FetchValidExchangeRate(date, sixMonths, currency string) (float64, error) {
	url := fmt.Sprintf(
		"https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=currency,exchange_rate,record_date&filter=currency:eq:%s,record_date:gte:%s,record_date:lte:%s&sort=-record_date",
		currency, sixMonths, date,
	)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("erro ao conectar à API: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Data []struct {
			ExchangeRate string `json:"exchange_rate"`
			RecordDate   string `json:"record_date"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	if len(data.Data) == 0 {
		return 0, fmt.Errorf("nenhuma taxa de câmbio encontrada para %s dentro dos últimos 6 meses", currency)
	}

	rate, err := strconv.ParseFloat(data.Data[0].ExchangeRate, 64)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter taxa de câmbio: %v", err)
	}

	return rate, nil
}
