package routes

import (
	"context"
	simplestorage "goledger-challenge-besu/abi"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, contract *simplestorage.SimpleStorage, client *ethclient.Client, auth *bind.TransactOpts) {
	router.POST("/set", func(ctx *gin.Context) {
		var req struct {
			Value int64 `json:"value"`
		}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
			return
		}
	
		tx, err := contract.Set(auth, big.NewInt(req.Value))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"status": "valor definido",
			"tx":     tx.Hash().Hex(),
		})
	})

	router.GET("/get", func(ctx *gin.Context) {
		value, err := contract.Get(&bind.CallOpts{Context: context.Background()})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"value": value.String()})
	})
}
