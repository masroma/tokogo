package responses

import "tokogo/models"

// UserManagementResponse struct untuk response user management
type UserManagementResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserListResponse struct untuk response list users
type UserListResponse struct {
	Users []UserManagementResponse `json:"users"`
	Total int                      `json:"total"`
	Page  int                      `json:"page"`
	Limit int                      `json:"limit"`
}

// ConvertUserToManagementResponse mengkonversi model User ke UserManagementResponse
func ConvertUserToManagementResponse(user models.User) UserManagementResponse {
	return UserManagementResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ConvertUsersToManagementResponse mengkonversi slice model User ke slice UserManagementResponse
func ConvertUsersToManagementResponse(users []models.User) []UserManagementResponse {
	var responses []UserManagementResponse
	for _, user := range users {
		responses = append(responses, ConvertUserToManagementResponse(user))
	}
	return responses
}
