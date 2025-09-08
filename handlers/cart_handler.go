package handlers

import (
	"net/http"
	"strconv"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService *services.CartService
}

func NewCartHandler() *CartHandler {
	return &CartHandler{
		cartService: services.NewCartService(),
	}
}

// AddToCart godoc
// @Summary Add item to cart
// @Description Add a product to user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart body requests.AddToCartRequest true "Cart item data"
// @Success 201 {object} responses.CartItemResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/cart [post]
func (h *CartHandler) AddToCart(c *gin.Context) {
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

	var req requests.AddToCartRequest

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

	// Call service to add to cart
	response, err := h.cartService.AddToCart(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "add_to_cart_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "Item added to cart successfully",
		Data:    response,
	})
}

// GetCart godoc
// @Summary Get user's cart
// @Description Get all items in user's cart
// @Tags Cart
// @Produce json
// @Success 200 {object} responses.CartResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/cart [get]
func (h *CartHandler) GetCart(c *gin.Context) {
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

	// Call service to get cart
	response, err := h.cartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_cart_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Cart retrieved successfully",
		Data:    response,
	})
}

// UpdateCartItem godoc
// @Summary Update cart item quantity
// @Description Update the quantity of a specific item in cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID"
// @Param cart body requests.UpdateCartItemRequest true "Updated quantity"
// @Success 200 {object} responses.CartItemResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/cart/{product_id} [put]
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
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

	// Get product ID from URL parameter
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid product ID",
		})
		return
	}

	var req requests.UpdateCartItemRequest

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

	// Call service to update cart item
	response, err := h.cartService.UpdateCartItem(userID, uint(productID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_cart_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Cart item updated successfully",
		Data:    response,
	})
}

// RemoveFromCart godoc
// @Summary Remove item from cart
// @Description Remove a specific item from user's cart
// @Tags Cart
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/cart/{product_id} [delete]
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
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

	// Get product ID from URL parameter
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid product ID",
		})
		return
	}

	// Call service to remove from cart
	err = h.cartService.RemoveFromCart(userID, uint(productID))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "remove_from_cart_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Item removed from cart successfully",
		Data:    nil,
	})
}

// ClearCart godoc
// @Summary Clear user's cart
// @Description Remove all items from user's cart
// @Tags Cart
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/cart/clear [delete]
func (h *CartHandler) ClearCart(c *gin.Context) {
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

	// Call service to clear cart
	err := h.cartService.ClearCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "clear_cart_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Cart cleared successfully",
		Data:    nil,
	})
}

// GetCartItemCount godoc
// @Summary Get cart item count
// @Description Get the total number of items in user's cart
// @Tags Cart
// @Produce json
// @Success 200 {object} map[string]int64
// @Failure 401 {object} map[string]string
// @Router /api/v1/cart/count [get]
func (h *CartHandler) GetCartItemCount(c *gin.Context) {
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

	// Call service to get cart count
	count, err := h.cartService.GetCartItemCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_cart_count_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Cart count retrieved successfully",
		Data:    gin.H{"count": count},
	})
}
