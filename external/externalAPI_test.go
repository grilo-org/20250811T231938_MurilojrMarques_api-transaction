package external_test

import (
	"testing"

	"github.com/MurilojrMarques/api-transaction.git/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFetchValidExchangeRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFetcher := mocks.NewMockExchangeRateFetcher(ctrl)

	date := "2023-12-01"
	sixMonths := "2023-06-01"
	currency := "Euro"

	mockFetcher.EXPECT().
		FetchValidExchangeRate(date, sixMonths, currency).
		Return(1.12, nil) // Simulando que a taxa de câmbio para Euro é 1.12

	result, err := mockFetcher.FetchValidExchangeRate(date, sixMonths, currency)

	assert.NoError(t, err)
	assert.Equal(t, 1.12, result)
}
