package services

import (
	"errors"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

// NewCategoryService membuat instance baru CategoryService
func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepo: repositories.NewCategoryRepository(),
	}
}

// CreateCategory membuat category baru
func (s *CategoryService) CreateCategory(req requests.CreateCategoryRequest) (*responses.CategoryResponse, error) {
	// Cek apakah nama category sudah ada
	exists, err := s.categoryRepo.CheckCategoryExists(req.Name, 0)
	if err != nil {
		return nil, errors.New("failed to check category existence")
	}
	if exists {
		return nil, errors.New("category name already exists")
	}

	// Buat category baru
	category := &models.Category{
		Name: req.Name,
	}

	// Simpan ke database
	if err := s.categoryRepo.CreateCategory(category); err != nil {
		return nil, errors.New("failed to create category")
	}

	// Return response
	return &responses.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetCategoryByID mengambil category berdasarkan ID
func (s *CategoryService) GetCategoryByID(id uint) (*responses.CategoryResponse, error) {
	category, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return &responses.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetAllCategories mengambil semua categories dengan pagination
func (s *CategoryService) GetAllCategories(page, limit int) (*responses.CategoryListResponse, error) {
	categories, total, err := s.categoryRepo.GetAllCategories(page, limit)
	if err != nil {
		return nil, errors.New("failed to get categories")
	}

	return &responses.CategoryListResponse{
		Categories: responses.ConvertCategoriesToResponse(categories),
		Total:      int(total),
		Page:       page,
		Limit:      limit,
	}, nil
}

// UpdateCategory mengupdate category berdasarkan ID
func (s *CategoryService) UpdateCategory(id uint, req requests.UpdateCategoryRequest) (*responses.CategoryResponse, error) {
	// Cek apakah category ada
	existingCategory, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	// Cek apakah nama category sudah ada (exclude current category)
	exists, err := s.categoryRepo.CheckCategoryExists(req.Name, id)
	if err != nil {
		return nil, errors.New("failed to check category existence")
	}
	if exists {
		return nil, errors.New("category name already exists")
	}

	// Update category
	existingCategory.Name = req.Name
	if err := s.categoryRepo.UpdateCategory(id, existingCategory); err != nil {
		return nil, errors.New("failed to update category")
	}

	// Get updated category
	updatedCategory, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, errors.New("failed to get updated category")
	}

	return &responses.CategoryResponse{
		ID:        updatedCategory.ID,
		Name:      updatedCategory.Name,
		Slug:      updatedCategory.Slug,
		CreatedAt: updatedCategory.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: updatedCategory.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// DeleteCategory menghapus category berdasarkan ID
func (s *CategoryService) DeleteCategory(id uint) error {
	// Cek apakah category ada
	_, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return err
	}

	// Hapus category
	if err := s.categoryRepo.DeleteCategory(id); err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}
