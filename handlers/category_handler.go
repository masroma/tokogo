package handlers

import (
	"net/http"
	"strconv"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

// NewCategoryHandler membuat instance baru CategoryHandler
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categoryService: services.NewCategoryService(),
	}
}

// CreateCategory handler untuk membuat category baru
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req requests.CreateCategoryRequest

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

	// Panggil service untuk create category
	categoryResponse, err := h.categoryService.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "create_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "Category created successfully",
		Data:    categoryResponse,
	})
}

// GetCategoryByID handler untuk mengambil category berdasarkan ID
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	// Ambil ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid category ID",
		})
		return
	}

	// Panggil service untuk get category
	categoryResponse, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Error:   "category_not_found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Category retrieved successfully",
		Data:    categoryResponse,
	})
}

// GetAllCategories handler untuk mengambil semua categories
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	// Ambil query parameters untuk pagination
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	// Panggil service untuk get all categories
	categoriesResponse, err := h.categoryService.GetAllCategories(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "get_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Categories retrieved successfully",
		Data:    categoriesResponse,
	})
}

// UpdateCategory handler untuk mengupdate category
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	// Ambil ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid category ID",
		})
		return
	}

	var req requests.UpdateCategoryRequest

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

	// Panggil service untuk update category
	categoryResponse, err := h.categoryService.UpdateCategory(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Category updated successfully",
		Data:    categoryResponse,
	})
}

// DeleteCategory handler untuk menghapus category
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	// Ambil ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid category ID",
		})
		return
	}

	// Panggil service untuk delete category
	err = h.categoryService.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "delete_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Category deleted successfully",
		Data:    nil,
	})
}
