package responses

import "tokogo/models"

// CategoryResponse struct untuk response category
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CategoryListResponse struct untuk response list category
type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Total      int                `json:"total"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
}

// ConvertCategoryToResponse mengkonversi Category model ke CategoryResponse
func ConvertCategoryToResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ConvertCategoriesToResponse mengkonversi slice Category ke slice CategoryResponse
func ConvertCategoriesToResponse(categories []models.Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, category := range categories {
		responses = append(responses, ConvertCategoryToResponse(category))
	}
	return responses
}
