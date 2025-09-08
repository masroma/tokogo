package handlers

import (
	"net/http"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

// NewProfileHandler membuat instance baru ProfileHandler
func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{
		profileService: services.NewProfileService(),
	}
}

// GetProfile handler untuk mengambil profile user
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	// Ambil user ID dari JWT token
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

	// Panggil service untuk get profile
	profile, err := h.profileService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "get_profile_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Profile retrieved successfully",
		Data:    profile,
	})
}

// UpdateProfile handler untuk mengupdate profile user
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	// Ambil user ID dari JWT token
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

	var req requests.UpdateProfileRequest

	// Bind dan validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Panggil service untuk update profile
	response, err := h.profileService.UpdateProfile(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_profile_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Profile updated successfully",
		Data:    response,
	})
}

// ChangePassword handler untuk mengubah password user
func (h *ProfileHandler) ChangeUserPassword(c *gin.Context) {
	// Ambil user ID dari JWT token
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

	var req requests.ChangeUserPasswordRequest

	// Bind dan validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Panggil service untuk change password
	response, err := h.profileService.ChangeUserPassword(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "change_password_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
