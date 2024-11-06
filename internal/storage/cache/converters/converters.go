package converters

import (
	"time"

	service_model "github.com/MGomed/auth/internal/model"
	cache_model "github.com/MGomed/auth/internal/storage/cache/model"
)

// ToUserInfoFromService converts service_model.UserInfo to cache_model.UserInfo
func ToUserInfoFromService(user *service_model.UserInfo) *cache_model.UserInfo {
	if user == nil {
		return nil
	}

	var updatedAt time.Time
	if user.UpdatedAt != nil {
		updatedAt = *user.UpdatedAt
	}

	return &cache_model.UserInfo{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.HashedPassword,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

// ToUserInfoFromCache converts cache_model.UserInfo to service_model.UserInfo
func ToUserInfoFromCache(user *cache_model.UserInfo) *service_model.UserInfo {
	if user == nil {
		return nil
	}

	var updatedAt *time.Time
	if !user.UpdatedAt.IsZero() {
		updatedAt = &user.UpdatedAt
	}

	return &service_model.UserInfo{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.Password,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      updatedAt,
	}
}
