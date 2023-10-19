package handler

import (
	"bwa/golang/helper"
	"bwa/golang/transaction"
	"bwa/golang/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionHandler struct {
	service transaction.TransactionService
}

func NewTransactionHandler(service transaction.TransactionService) *transactionHandler {
	return &transactionHandler{
		service: service,
	}
}
func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)

	fmt.Println("nilai id adalah : ", input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	transactions, err := h.service.GetTransactionByCampaignId(input)

	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transactions dari service", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign's Transaction", http.StatusOK, "success", transaction.TransactionsFormatter(transactions))
	c.JSON(http.StatusOK, response)
}
