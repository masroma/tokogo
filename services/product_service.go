package services

import (
	"errors"
	"tokogo/config"
	"tokogo/helpers"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type ProductService struct {
	productRepo  *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

// NewProductService membuat instance baru ProductService
func NewProductService() *ProductService {
	return &ProductService{
		productRepo:  repositories.NewProductRepository(config.DB),
		categoryRepo: repositories.NewCategoryRepository(),
	}
}

// CreateProduct membuat product baru
func (s *ProductService) CreateProduct(req requests.CreateProductRequest, imagePath string) (*responses.ProductResponse, error) {
	// Cek apakah category ada
	_, err := s.categoryRepo.GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Buat product baru
	product := &models.Product{
		Name:          req.Name,
		Description:   req.Description,
		PurchasePrice: req.PurchasePrice,
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		CategoryID:    req.CategoryID,
		ImageURL:      imagePath,
	}

	// Simpan ke database
	if err := s.productRepo.Create(product); err != nil {
		// If database save fails, delete uploaded file
		if imagePath != "" {
			helpers.DeleteFile(imagePath)
		}
		return nil, errors.New("failed to create product")
	}

	// Return response
	response := responses.ConvertProductToResponse(*product)
	return &response, nil
}

func (s *ProductService) GetAllProducts(page, limit int) (*responses.ProductListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := s.productRepo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	productResponses := responses.ConvertProductsToResponse(products)

	return &responses.ProductListResponse{
		Products: productResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *ProductService) GetProductByID(id uint) (*responses.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := responses.ConvertProductToResponse(*product)
	return &response, nil
}

func (s *ProductService) UpdateProduct(id uint, req requests.UpdateProductRequest) (*responses.ProductResponse, error) {
	// Get existing product
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Validate category exists
	_, err = s.categoryRepo.GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Update product fields
	product.Name = req.Name
	product.Description = req.Description
	product.PurchasePrice = req.PurchasePrice
	product.SellingPrice = req.SellingPrice
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	// ImageURL is not updated via request - handled separately

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	response := responses.ConvertProductToResponse(*product)
	return &response, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(id)
}

func (s *ProductService) GetProductsByCategory(categoryID uint, page, limit int) (*responses.ProductListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := s.productRepo.GetByCategoryID(categoryID, page, limit)
	if err != nil {
		return nil, err
	}

	productResponses := responses.ConvertProductsToResponse(products)

	return &responses.ProductListResponse{
		Products: productResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

// Public methods (tanpa purchase_price)
func (s *ProductService) GetAllProductsPublic(page, limit int) (*responses.PublicProductListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := s.productRepo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	productResponses := responses.ConvertProductsToPublicResponse(products)

	return &responses.PublicProductListResponse{
		Products: productResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *ProductService) GetProductByIDPublic(id uint) (*responses.PublicProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := responses.ConvertProductToPublicResponse(*product)
	return &response, nil
}

func (s *ProductService) GetProductsByCategoryPublic(categoryID uint, page, limit int) (*responses.PublicProductListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := s.productRepo.GetByCategoryID(categoryID, page, limit)
	if err != nil {
		return nil, err
	}

	productResponses := responses.ConvertProductsToPublicResponse(products)

	return &responses.PublicProductListResponse{
		Products: productResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
