package responses

import "tokogo/models"

// ProfileResponse struct untuk response profile
type ProfileResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ChangeUserPasswordResponse struct untuk response change password
type ChangeUserPasswordResponse struct {
	Message string `json:"message"`
}

// ConvertUserToProfileResponse mengkonversi model User ke ProfileResponse
func ConvertUserToProfileResponse(user models.User) ProfileResponse {
	return ProfileResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
