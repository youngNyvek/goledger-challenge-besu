package controller

import (
	"goledger-challenge-besu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	service *service.ContractService
}

func NewContractController(svc *service.ContractService) *ContractController {
	return &ContractController{service: svc}
}

func (c *ContractController) SetValue(ctx *gin.Context) {
	var req struct { Value int64 `json:"value"` }
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	txHash, blockNumber, err := c.service.SetValue(ctx.Request.Context(), req.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Transaction confirmed",
		"txHash":      txHash,
		"blockNumber": blockNumber,
	})
}

func (c *ContractController) GetValue(ctx *gin.Context) {
	val, err := c.service.GetValue(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"value": val})
}

func (c *ContractController) Sync(ctx *gin.Context) {
	val, err := c.service.Sync(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"syncedValue": val})
}

func (c *ContractController) Check(ctx *gin.Context) {
	equal, err := c.service.Check(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ "isEqual": equal	})
}