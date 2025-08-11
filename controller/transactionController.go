package controller

import (
	"net/http"
	"strconv"

	"github.com/MurilojrMarques/api-transaction.git/model"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(usecase usecase.TransactionUsecase) transactionController {
	return transactionController{
		transactionUsecase: usecase,
	}
}

// @BasePath api/v1

// CreateTransaction godoc
// @Summary Cria uma nova transação
// @Description Cria uma nova transação de compra no banco de dados com os dados fornecidos
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body model.Transaction true "Dados da transação"
// @Success 201 {object} model.Transaction
// @Failure 400 {object} model.ErrorResponse "Dados inválidos"
// @Failure 500 {object} model.ErrorResponse "Erro interno no servidor"
// @Router /transaction [post]
func (t *transactionController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	err := ctx.BindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Dados inválidos"})
		return
	}

	insertedTransaction, err := t.transactionUsecase.CreateTransaction(transaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erros": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedTransaction)
}

// GetTransactionConverted godoc
// @Summary Retorna uma transação convertida para outra moeda
// @Description Busca uma transação pelo ID e converte seu valor para a moeda especificada
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param id path int true "ID da transação"
// @Param currency query string false "Moeda para conversão (padrão: USD)"
// @Success 200 {object} model.ConvertedTransaction
// @Failure 400 {object} model.ErrorResponse "ID inválido"
// @Failure 500 {object} model.ErrorResponse "Erro interno no servidor"
// @Router /transaction/{id}/convert [get]
func (t *transactionController) GetTransactionConverted(ctx *gin.Context) {
	id := ctx.Param("id")
	currency := ctx.DefaultQuery("currency", "USD")

	transactionID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido. Deve ser um número inteiro"})
		return
	}

	transaction, err := t.transactionUsecase.GetTransactionConverted(transactionID, currency)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
