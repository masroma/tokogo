package handlers

import (
	"net/http"
	"strconv"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type UserManagementHandler struct {
	userService *services.UserManagementService
}

// NewUserManagementHandler membuat instance baru UserManagementHandler
func NewUserManagementHandler() *UserManagementHandler {
	return &UserManagementHandler{
		userService: services.NewUserManagementService(),
	}
}

// CreateUser handler untuk membuat user baru
func (h *UserManagementHandler) CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest

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

	// Panggil service untuk create user
	userResponse, err := h.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "create_user_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "User created successfully",
		Data:    userResponse,
	})
}

// GetUserByID handler untuk mengambil user berdasarkan ID
func (h *UserManagementHandler) GetUserByID(c *gin.Context) {
	// Ambil ID dari parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid user ID",
		})
		return
	}

	// Panggil service untuk get user
	userResponse, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Error:   "user_not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "User retrieved successfully",
		Data:    userResponse,
	})
}

// GetAllUsers handler untuk mengambil semua users
func (h *UserManagementHandler) GetAllUsers(c *gin.Context) {
	// Ambil parameter pagination
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	// Panggil service untuk get all users
	usersResponse, err := h.userService.GetAllUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_users_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Users retrieved successfully",
		Data:    usersResponse,
	})
}

// UpdateUser handler untuk mengupdate user
func (h *UserManagementHandler) UpdateUser(c *gin.Context) {
	// Ambil ID dari parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid user ID",
		})
		return
	}

	var req requests.UpdateUserRequest

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

	// Panggil service untuk update user
	userResponse, err := h.userService.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_user_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "User updated successfully",
		Data:    userResponse,
	})
}

// DeleteUser handler untuk menghapus user
func (h *UserManagementHandler) DeleteUser(c *gin.Context) {
	// Ambil ID dari parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid user ID",
		})
		return
	}

	// Panggil service untuk delete user
	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "delete_user_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "User deleted successfully",
		Data:    nil,
	})
}

// UpdateUserRole handler untuk mengupdate role user
func (h *UserManagementHandler) UpdateUserRole(c *gin.Context) {
	// Ambil ID dari parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid user ID",
		})
		return
	}

	var req struct {
		Role string `json:"role" binding:"required,oneof=customer admin"`
	}

	// Bind dan validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Panggil service untuk update user role
	userResponse, err := h.userService.UpdateUserRole(uint(id), req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_user_role_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "User role updated successfully",
		Data:    userResponse,
	})
}

// GetUsersByRole handler untuk mengambil users berdasarkan role
func (h *UserManagementHandler) GetUsersByRole(c *gin.Context) {
	// Ambil role dari parameter
	role := c.Param("role")
	if role != "customer" && role != "admin" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_role",
			Message: "Role must be either customer or admin",
		})
		return
	}

	// Ambil parameter pagination
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	// Panggil service untuk get users by role
	usersResponse, err := h.userService.GetUsersByRole(role, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_users_by_role_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Users retrieved successfully",
		Data:    usersResponse,
	})
}
