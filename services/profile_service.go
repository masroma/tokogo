package services

import (
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type ProfileService struct {
	profileRepo *repositories.ProfileRepository
}

// NewProfileService membuat instance baru ProfileService
func NewProfileService() *ProfileService {
	return &ProfileService{
		profileRepo: repositories.NewProfileRepository(),
	}
}

// GetProfile mengambil profile user
func (s *ProfileService) GetProfile(userID uint) (*responses.ProfileResponse, error) {
	user, err := s.profileRepo.GetProfileByID(userID)
	if err != nil {
		return nil, err
	}

	profileResponse := responses.ConvertUserToProfileResponse(*user)
	return &profileResponse, nil
}

// UpdateProfile mengupdate profile user
func (s *ProfileService) UpdateProfile(userID uint, req requests.UpdateProfileRequest) (*responses.ProfileResponse, error) {
	// Validasi request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Update profile
	user, err := s.profileRepo.UpdateProfile(userID, req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	profileResponse := responses.ConvertUserToProfileResponse(*user)
	return &profileResponse, nil
}

// ChangePassword mengubah password user
func (s *ProfileService) ChangeUserPassword(userID uint, req requests.ChangeUserPasswordRequest) (*responses.ChangeUserPasswordResponse, error) {
	// Validasi request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Change password
	if err := s.profileRepo.ChangeUserPassword(userID, req.CurrentPassword, req.NewPassword); err != nil {
		return nil, err
	}

	return &responses.ChangeUserPasswordResponse{
		Message: "Password changed successfully",
	}, nil
}
