package usecase_test

import (
	"testing"
	"time"

	"github.com/MurilojrMarques/api-transaction.git/model"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(transaction model.Transaction) (int, error) {
	args := m.Called(transaction)
	return args.Int(0), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionByID(id int) (model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepository)

	mockRepo.On("CreateTransaction", mock.Anything).Return(1, nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockRepo)

	inputTransaction := model.Transaction{
		Description: "Test Transaction",
		Value:       100.50,
		Date:        model.CustomDate(time.Now().Truncate(24 * time.Hour)),
	}

	result, err := transactionUsecase.CreateTransaction(inputTransaction)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID) // Esperando que o ID da transação seja 1
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionConverted(t *testing.T) {
	mockRepo := new(MockTransactionRepository)

	transaction := model.Transaction{
		ID:          1,
		Description: "Test Transaction",
		Value:       100.50,
		Date:        model.CustomDate(time.Now().Truncate(24 * time.Hour)),
	}
	mockRepo.On("GetTransactionByID", 1).Return(transaction, nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockRepo)

	result, err := transactionUsecase.GetTransactionConverted(1, "Euro")

	assert.NoError(t, err)
	assert.Equal(t, 1, result.Transaction.ID)
	mockRepo.AssertExpectations(t)
}
