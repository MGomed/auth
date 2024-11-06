package converters

import (
	service_model "github.com/MGomed/auth/internal/model"
	api "github.com/MGomed/auth/pkg/user_api"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserInfoFromService converts UserInfo to api.UserInfo
func ToUserInfoFromService(user *service_model.UserInfo) *api.UserInfo {
	if user == nil {
		return nil
	}

	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &api.UserInfo{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      api.Role(api.Role_value[user.Role]),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserCreateFromAPI converts api.UserCreate to UserCreate
func ToUserCreateFromAPI(user *api.UserCreate) *service_model.UserCreate {
	if user == nil {
		return nil
	}

	return &service_model.UserCreate{
		Name:     user.Name,
		Email:    user.Email,
		Password: []byte(user.Password),
		Role:     service_model.RoleNames[int32(user.Role)],
	}
}

// ToUserUpdateFromAPI converts api.UserUpdate to UserUpdate
func ToUserUpdateFromAPI(user *api.UserUpdate) *service_model.UserUpdate {
	if user == nil {
		return nil
	}

	var res service_model.UserUpdate

	res.ID = user.Id

	if val := user.Name.GetValue(); val != "" {
		res.Name = &val
	}

	role := api.Role_name[int32(user.Role)]
	res.Role = &role

	return &res
}
