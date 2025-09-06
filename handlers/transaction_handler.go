package handlers

import (
	"net/http"
	"strconv"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

// NewTransactionHandler membuat instance baru TransactionHandler
func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{
		transactionService: services.NewTransactionService(),
	}
}

// GetAllTransactions handler untuk mengambil semua transaksi
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	// Parse query parameters
	status := c.Query("status")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid page parameter",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid limit parameter",
		})
		return
	}

	// Get transactions
	result, err := h.transactionService.GetAllTransactions(page, limit, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_transactions_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Transactions retrieved successfully",
		Data:    result,
	})
}

// GetTransactionByID handler untuk mengambil transaksi berdasarkan ID
func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	// Parse transaction ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid transaction ID",
		})
		return
	}

	// Get transaction
	transaction, err := h.transactionService.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Error:   "transaction_not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	})
}

// UpdateTransactionStatus handler untuk mengupdate status transaksi
func (h *TransactionHandler) UpdateTransactionStatus(c *gin.Context) {
	// Parse transaction ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid transaction ID",
		})
		return
	}

	// Parse request body
	var req requests.UpdateTransactionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Update transaction status
	transaction, err := h.transactionService.UpdateTransactionStatus(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_transaction_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Transaction status updated successfully",
		Data:    transaction,
	})
}
