package services

import (
	"errors"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"

	"golang.org/x/crypto/bcrypt"
)

type UserManagementService struct {
	userRepo *repositories.UserManagementRepository
}

// NewUserManagementService membuat instance baru UserManagementService
func NewUserManagementService() *UserManagementService {
	return &UserManagementService{
		userRepo: repositories.NewUserManagementRepository(),
	}
}

// CreateUser membuat user baru
func (s *UserManagementService) CreateUser(req requests.CreateUserRequest) (*responses.UserManagementResponse, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Buat user baru
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	// Simpan ke database
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Return response
	response := responses.ConvertUserToManagementResponse(*user)
	return &response, nil
}

// GetUserByID mengambil user berdasarkan ID
func (s *UserManagementService) GetUserByID(id uint) (*responses.UserManagementResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	response := responses.ConvertUserToManagementResponse(*user)
	return &response, nil
}

// GetAllUsers mengambil semua users dengan pagination
func (s *UserManagementService) GetAllUsers(page, limit int) (*responses.UserListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := s.userRepo.GetAllUsers(page, limit)
	if err != nil {
		return nil, errors.New("failed to get users")
	}

	return &responses.UserListResponse{
		Users: responses.ConvertUsersToManagementResponse(users),
		Total: int(total),
		Page:  page,
		Limit: limit,
	}, nil
}

// UpdateUser mengupdate user
func (s *UserManagementService) UpdateUser(id uint, req requests.UpdateUserRequest) (*responses.UserManagementResponse, error) {
	// Ambil user yang akan diupdate
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update field yang ada
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Cek apakah email sudah digunakan user lain
		existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already used by another user")
		}
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}

	// Simpan perubahan
	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	// Return response
	response := responses.ConvertUserToManagementResponse(*user)
	return &response, nil
}

// DeleteUser menghapus user
func (s *UserManagementService) DeleteUser(id uint) error {
	// Cek apakah user ada
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Hapus user
	if err := s.userRepo.DeleteUser(id); err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

// UpdateUserRole mengupdate role user
func (s *UserManagementService) UpdateUserRole(id uint, role string) (*responses.UserManagementResponse, error) {
	// Cek apakah user ada
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update role
	if err := s.userRepo.UpdateUserRole(id, role); err != nil {
		return nil, errors.New("failed to update user role")
	}

	// Ambil user yang sudah diupdate
	updatedUser, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("failed to get updated user")
	}

	// Return response
	response := responses.ConvertUserToManagementResponse(*updatedUser)
	return &response, nil
}

// GetUsersByRole mengambil users berdasarkan role
func (s *UserManagementService) GetUsersByRole(role string, page, limit int) (*responses.UserListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	if role != "customer" && role != "admin" {
		return nil, errors.New("invalid role")
	}

	users, total, err := s.userRepo.GetUsersByRole(role, page, limit)
	if err != nil {
		return nil, errors.New("failed to get users by role")
	}

	return &responses.UserListResponse{
		Users: responses.ConvertUsersToManagementResponse(users),
		Total: int(total),
		Page:  page,
		Limit: limit,
	}, nil
}
