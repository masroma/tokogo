package handlers

import (
	"net/http"
	"strconv"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	checkoutService *services.CheckoutService
}

func NewCheckoutHandler() *CheckoutHandler {
	return &CheckoutHandler{
		checkoutService: services.NewCheckoutService(),
	}
}

// GetCheckoutSummary godoc
// @Summary Get checkout summary
// @Description Get checkout summary with shipping cost calculation
// @Tags Checkout
// @Accept json
// @Produce json
// @Param checkout body requests.CheckoutRequest true "Checkout data"
// @Success 200 {object} responses.CheckoutSummaryResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/checkout/summary [post]
func (h *CheckoutHandler) GetCheckoutSummary(c *gin.Context) {
	// Get user ID from JWT token
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User ID not found",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid user ID",
		})
		return
	}

	var req requests.CheckoutRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validate using method Validate()
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Call service to get checkout summary
	response, err := h.checkoutService.GetCheckoutSummary(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "get_checkout_summary_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Checkout summary retrieved successfully",
		Data:    response,
	})
}

// ProcessCheckout godoc
// @Summary Process checkout
// @Description Process checkout from cart to transaction
// @Tags Checkout
// @Accept json
// @Produce json
// @Param checkout body requests.CheckoutRequest true "Checkout data"
// @Success 201 {object} responses.CheckoutResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/checkout [post]
func (h *CheckoutHandler) ProcessCheckout(c *gin.Context) {
	// Get user ID from JWT token
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User ID not found",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid user ID",
		})
		return
	}

	var req requests.CheckoutRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validate using method Validate()
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Call service to process checkout
	response, err := h.checkoutService.ProcessCheckout(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "checkout_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "Checkout processed successfully",
		Data:    response,
	})
}

// ConfirmPayment godoc
// @Summary Confirm payment
// @Description Confirm payment for a transaction
// @Tags Checkout
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Param payment body requests.ConfirmPaymentRequest true "Payment confirmation data"
// @Success 200 {object} responses.CheckoutResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/checkout/{transaction_id}/confirm [post]
func (h *CheckoutHandler) ConfirmPayment(c *gin.Context) {
	// Get user ID from JWT token
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User ID not found",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid user ID",
		})
		return
	}

	// Get transaction ID from URL parameter
	transactionID, err := strconv.ParseUint(c.Param("transaction_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid transaction ID",
		})
		return
	}

	var req requests.ConfirmPaymentRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validate using method Validate()
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Call service to confirm payment
	response, err := h.checkoutService.ConfirmPayment(userID, uint(transactionID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "confirm_payment_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Payment confirmed successfully",
		Data:    response,
	})
}

// GetUserTransactions godoc
// @Summary Get user transactions
// @Description Get all transactions for the authenticated user
// @Tags Checkout
// @Produce json
// @Success 200 {object} []responses.CheckoutResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/checkout/transactions [get]
func (h *CheckoutHandler) GetUserTransactions(c *gin.Context) {
	// Get user ID from JWT token
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User ID not found",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid user ID",
		})
		return
	}

	// Call service to get user transactions
	response, err := h.checkoutService.GetUserTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_transactions_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Transactions retrieved successfully",
		Data:    response,
	})
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Description Get a specific transaction by ID for the authenticated user
// @Tags Checkout
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Success 200 {object} responses.CheckoutResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/checkout/transactions/{transaction_id} [get]
func (h *CheckoutHandler) GetTransactionByID(c *gin.Context) {
	// Get user ID from JWT token
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User ID not found",
		})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid user ID",
		})
		return
	}

	// Get transaction ID from URL parameter
	transactionID, err := strconv.ParseUint(c.Param("transaction_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid transaction ID",
		})
		return
	}

	// Call service to get transaction
	response, err := h.checkoutService.GetTransactionByID(userID, uint(transactionID))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "get_transaction_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Transaction retrieved successfully",
		Data:    response,
	})
}
