package handlers

import (
	"net/http"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler membuat instance baru AuthHandler
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

// Register handler untuk registrasi user baru
func (h *AuthHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest

	// Bind dan validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validasi menggunakan method Validate()
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Panggil service untuk register
	registerResponse, err := h.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "register_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "User registered successfully",
		Data:    registerResponse,
	})
}

// Login handler untuk login user
func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.LoginRequest

	// Bind dan validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validasi menggunakan method Validate()
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Panggil service untuk login
	loginResponse, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "login_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Login successful",
		Data:    loginResponse,
	})
}

// Logout handler untuk logout user
func (h *AuthHandler) Logout(c *gin.Context) {
	// Ambil user ID dari context (setelah AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User not authenticated",
		})
		return
	}

	// Panggil service untuk logout
	logoutResponse, err := h.authService.Logout(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "logout_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, logoutResponse)
}

// GetProfile handler untuk mengambil profile user
func (h *AuthHandler) GetProfile(c *gin.Context) {
	// Ambil user ID dari context (setelah AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error:   "unauthorized",
			Message: "User not authenticated",
		})
		return
	}

	// Panggil service untuk get profile
	userResponse, err := h.authService.GetProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Error:   "user_not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Profile retrieved successfully",
		Data:    userResponse,
	})
}
