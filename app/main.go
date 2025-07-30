package main

import (
	"fmt"
	contract_controller "goledger-challenge-besu/internal/controllers"
	"goledger-challenge-besu/internal/infra/database"
	ethereum "goledger-challenge-besu/internal/infra/ehtereum"
	"goledger-challenge-besu/internal/repository"
	"goledger-challenge-besu/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	client := ethereum.MustConnectToNode("http://localhost:8545")
	contract := ethereum.MustCreateContract("0x42699A7612A82f1d9C36148af9C77354759b210b", client)
	signer := ethereum.MustCreateSigner("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63", client)
	db := database.MustConnectPostgres()
	repo := repository.NewPostgresContractLogRepository(db)
	svc := service.NewContractService(contract, client, signer, repo)
	
	router := gin.Default()
	ctrl := contract_controller.NewContractController(svc)
	
	api := router.Group("/api/v1/contract")
	{
			api.POST("/set",   ctrl.SetValue)
			api.GET("/get",    ctrl.GetValue)
			api.POST("/sync",  ctrl.Sync)
			api.GET("/check",  ctrl.Check)
	}

	router.Run(":8080")
	fmt.Println("Servidor iniciado em :8080")
}
