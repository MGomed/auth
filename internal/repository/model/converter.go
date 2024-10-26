package model

import (
	service_model "github.com/MGomed/auth/internal/model"
)

// ToUserFromRepo converts repo_model.UserInfo to service_model.UserInfo
func ToUserFromRepo(user *UserInfo) *service_model.UserInfo {
	var updatedAt = &user.UpdatedAt.Time
	if !user.UpdatedAt.Valid {
		updatedAt = nil
	}

	return &service_model.UserInfo{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.Password,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      updatedAt,
	}
}
