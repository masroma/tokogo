package handlers

import (
	"net/http"
	"strconv"
	"tokogo/helpers"
	"tokogo/requests"
	"tokogo/responses"
	"tokogo/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler membuat instance baru ProductHandler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: services.NewProductService(),
	}
}

// CreateProduct handler untuk membuat product baru
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req requests.CreateProductRequest

	// Bind form data (including file upload)
	if err := c.ShouldBind(&req); err != nil {
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

	// Handle file upload
	var imagePath string
	if file, err := c.FormFile("image"); err == nil {
		// File was uploaded
		uploadDir := "./uploads/products"
		uploadedPath, err := helpers.UploadFile(file, uploadDir)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error:   "upload_failed",
				Message: err.Error(),
			})
			return
		}
		imagePath = uploadedPath
	}

	// Panggil service untuk create product
	productResponse, err := h.productService.CreateProduct(req, imagePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error:   "create_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
		Message: "Product created successfully",
		Data:    productResponse,
	})
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products with pagination (Admin only)
// @Tags Products
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} responses.ProductListResponse
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /admin/products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, err := h.productService.GetAllProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a specific product by ID (Admin only)
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update a product by ID (Admin only)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body requests.UpdateProductRequest true "Product data"
// @Success 200 {object} responses.ProductResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req requests.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product by ID (Admin only)
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// GetProductsByCategory godoc
// @Summary Get products by category
// @Description Get products by category ID with pagination (Admin only)
// @Tags Products
// @Produce json
// @Param category_id path int true "Category ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} responses.ProductListResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /admin/categories/{category_id}/products [get]
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("category_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, err := h.productService.GetProductsByCategory(uint(categoryID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
