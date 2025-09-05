package responses

import "tokogo/models"

// RegisterResponse struct untuk response register
type RegisterResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

// UserResponse struct untuk response user (tanpa password)
type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

// ErrorResponse struct untuk response error
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SuccessResponse struct untuk response sukses
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ConvertUserToResponse mengkonversi User model ke UserResponse
func ConvertUserToResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
