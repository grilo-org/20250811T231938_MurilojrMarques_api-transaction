package usecase

import (
	"fmt"
	"math"
	"time"

	"github.com/MurilojrMarques/api-transaction.git/external"
	"github.com/MurilojrMarques/api-transaction.git/model"
	"github.com/MurilojrMarques/api-transaction.git/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionUsecase struct {
	repository repository.TransactionRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{
		repository: repo,
	}
}

func (tu *TransactionUsecase) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	validate := validator.New()
	if err := validate.Struct(transaction); err != nil {
		return model.Transaction{}, fmt.Errorf("campo nulo ou valor menor que zero. erro: %v", err)
	}

	if len(transaction.Description) > 50 {
		return model.Transaction{}, fmt.Errorf("a descrição não pode exceder 50 caracteres")
	}

	transaction.Value = math.Round(transaction.Value*100) / 100

	transactionId, err := tu.repository.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("erro ao criar transação: %s", err.Error())
	}

	transaction.ID = transactionId

	return transaction, nil
}

func (tu *TransactionUsecase) GetTransactionConverted(transactionID int, currency string) (model.ConvertedTransaction, error) {
	transaction, err := tu.repository.GetTransactionByID(transactionID)
	if err != nil {
		return model.ConvertedTransaction{}, fmt.Errorf("transação não encontrada: %v", err)
	}

	dateFormatted := time.Time(transaction.Date).Format("2006-01-02")
	sixMonthsBack := time.Time(transaction.Date).AddDate(0, -6, 0).Format("2006-01-02")

	exchangeRate, err := external.FetchValidExchangeRate(dateFormatted, sixMonthsBack, currency)
	if err != nil {
		return model.ConvertedTransaction{}, fmt.Errorf("erro ao obter taxa de câmbio: %v", err)
	}

	// Calcular o valor convertido
	convertedValue := math.Round(transaction.Value*exchangeRate*100) / 100

	return model.ConvertedTransaction{
		Transaction:    transaction,
		ExchangeRate:   exchangeRate,
		ConvertedValue: convertedValue,
	}, nil
}
