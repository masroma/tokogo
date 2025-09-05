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
