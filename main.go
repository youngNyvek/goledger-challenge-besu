package main

import (
	"fmt"
	"goledger-challenge-besu/internal/blockchain"
	"goledger-challenge-besu/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	client := blockchain.MustConnectToNode("http://localhost:8545")

	contract := blockchain.MustCreateContract("0x42699A7612A82f1d9C36148af9C77354759b210b", client)

	transactor := blockchain.MustCreateTransactor("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63", client)

	router := gin.Default()
	routes.RegisterRoutes(router, contract, client, transactor)

	fmt.Println("Servidor iniciado em :8080")
	router.Run(":8080")
}
