package main

import (
	"log"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/MurilojrMarques/api-transaction.git/controller"
	"github.com/MurilojrMarques/api-transaction.git/database/config"
	"github.com/MurilojrMarques/api-transaction.git/docs"
	"github.com/MurilojrMarques/api-transaction.git/repository"
	"github.com/MurilojrMarques/api-transaction.git/usecase"
	"github.com/gin-gonic/gin"
)

// @title  Transaction api
// @version 1.0
// @description Api para conversão de transação
// @termsOfService http://swagger.io/terms/
func main() {
	//banco de dados
	postgresDb, err := config.NewPostgresDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer postgresDb.Db.Close()

	//camadas
	TransactionRepositoryImplementation := repository.NewTransactionRepository(postgresDb.Db)
	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepositoryImplementation)
	TransactionController := controller.NewTransactionController(*TransactionUsecase)

	//routes
	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := server.Group("/api/v1")
	{
		v1.POST("/transaction", TransactionController.CreateTransaction)
		v1.GET("/transaction/:id/convert", TransactionController.GetTransactionConverted)
	}
	//Iniciar o swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//servidor
	if err := server.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
	}
	log.Println("Servidor rodando na porta 8080")

}
