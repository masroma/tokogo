package responses

import "tokogo/models"

type ProductResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
	CategoryID    uint    `json:"category_id"`
	CategoryName  string  `json:"category_name"`
	ImagePath     string  `json:"image_url"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

func ConvertProductToResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		PurchasePrice: product.PurchasePrice,
		SellingPrice:  product.SellingPrice,
		Stock:         product.Stock,
		CategoryID:    product.CategoryID,
		CategoryName:  product.Category.Name,
		ImagePath:     product.ImageURL,
		CreatedAt:     product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertProductsToResponse(products []models.Product) []ProductResponse {
	var responses []ProductResponse
	for _, product := range products {
		responses = append(responses, ConvertProductToResponse(product))
	}
	return responses
}

// PublicProductResponse untuk response public (tanpa purchase_price)
type PublicProductResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	SellingPrice float64 `json:"selling_price"`
	Stock        int     `json:"stock"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	ImagePath    string  `json:"image_url"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type PublicProductListResponse struct {
	Products []PublicProductResponse `json:"products"`
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	Limit    int                     `json:"limit"`
}

func ConvertProductToPublicResponse(product models.Product) PublicProductResponse {
	return PublicProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		SellingPrice: product.SellingPrice,
		Stock:        product.Stock,
		CategoryID:   product.CategoryID,
		CategoryName: product.Category.Name,
		ImagePath:    product.ImageURL,
		CreatedAt:    product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertProductsToPublicResponse(products []models.Product) []PublicProductResponse {
	var responses []PublicProductResponse
	for _, product := range products {
		responses = append(responses, ConvertProductToPublicResponse(product))
	}
	return responses
}
