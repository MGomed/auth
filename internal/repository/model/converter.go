package model

import (
	service_model "github.com/MGomed/auth/internal/model"
)

func ToUserFromRepo(user *User) *service_model.User {
	var updatedAt = &user.UpdatedAt.Time
	if !user.UpdatedAt.Valid {
		updatedAt = nil
	}

	return &service_model.User{
		ID:        &user.ID,
		Name:      &user.Name,
		Email:     &user.Email,
		Password:  &user.Password,
		Role:      &user.Role,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
