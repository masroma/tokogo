package services

import (
	"errors"
	"tokogo/helpers"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

// NewAuthService membuat instance baru AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		authRepo: repositories.NewAuthRepository(),
	}
}

// Register mendaftarkan user baru
func (s *AuthService) Register(req requests.RegisterRequest) (*responses.RegisterResponse, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.authRepo.GetUserByEmail(req.Email)
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
		Name:     req.Username, // Menggunakan Username dari request
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "customer", // Default role
	}

	// Simpan ke database
	if err := s.authRepo.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Return response
	return &responses.RegisterResponse{
		User:  responses.ConvertUserToResponse(*user),
		Token: token,
	}, nil
}

// Login melakukan login user
func (s *AuthService) Login(req requests.LoginRequest) (*responses.LoginResponse, error) {
	// Cari user berdasarkan email
	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Return response
	return &responses.LoginResponse{
		User:  responses.ConvertUserToResponse(*user),
		Token: token,
	}, nil
}
