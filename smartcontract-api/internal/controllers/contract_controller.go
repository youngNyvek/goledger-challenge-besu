package contract_controller

import (
	"context"
	simplestorage "goledger-challenge-besu/abi"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)


type ContractController struct {
	contract *simplestorage.SimpleStorage
	client   *ethclient.Client
	signer   *bind.TransactOpts
}

func NewContractController(
	contract *simplestorage.SimpleStorage,
	client *ethclient.Client,
	signer *bind.TransactOpts,
) *ContractController {
	return &ContractController{
		contract: contract,
		client:   client,
		signer:   signer,
	}
}

func (c *ContractController) SetValue(ctx *gin.Context) {
	var request struct {
		Value int64 `json:"value"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	tx, err := c.contract.Set(c.signer, big.NewInt(request.Value))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Valor enviado com sucesso", "tx": tx.Hash().Hex()})
}

func (c *ContractController) GetValue(ctx *gin.Context) {
	value, err := c.contract.Get(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"value": value.String()})
}
