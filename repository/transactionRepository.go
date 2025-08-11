package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MurilojrMarques/api-transaction.git/model"
)

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) (int, error)
	GetTransactionByID(id int) (model.Transaction, error)
}

type TransactionRepositoryImplementation struct {
	connection *sql.DB
}

func NewTransactionRepository(connection *sql.DB) TransactionRepository {
	return &TransactionRepositoryImplementation{
		connection: connection,
	}
}

func (tr *TransactionRepositoryImplementation) CreateTransaction(transaction model.Transaction) (int, error) {

	var id int
	query, err := tr.connection.Prepare("INSERT INTO transaction" +
		"(description, date, value)" +
		"VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(transaction.Description, time.Time(transaction.Date), transaction.Value).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (tr *TransactionRepositoryImplementation) GetTransactionByID(id int) (model.Transaction, error) {
	var transaction model.Transaction
	query := "SELECT id, description, date, value FROM transaction WHERE id = $1"
	err := tr.connection.QueryRow(query, id).Scan(&transaction.ID, &transaction.Description, &transaction.Date, &transaction.Value)
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}
