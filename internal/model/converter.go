package model

import (
	api "github.com/MGomed/auth/pkg/user_api"
	"google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromService(user *User) *api.User {
	var res api.User
	if user.ID != nil {
		res.Id = wrapperspb.Int64(*user.ID)
	}

	if user.Name != nil {
		res.Name = wrapperspb.String(*user.Name)
	}

	if user.Email != nil {
		res.Email = wrapperspb.String(*user.Email)
	}

	if user.Password != nil {
		res.Password = wrapperspb.String(*user.Password)
	}

	if user.Role != nil {
		res.Role = api.Role(api.Role_value[*user.Role])
	}

	if user.CreatedAt != nil {
		res.CreatedAt = timestamppb.New(*user.CreatedAt)
	}

	if user.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &res
}

func ToUserFromApi(user *api.User) *User {
	var res User
	if user.Id != nil {
		res.ID = &user.Id.Value
	}

	if user.Name != nil {
		res.Name = &user.Name.Value
	}

	if user.Email != nil {
		res.Email = &user.Email.Value
	}

	if user.Password != nil {
		res.Password = &user.Password.Value
	}

	role := api.Role_name[int32(user.Role)]
	res.Role = &role

	if user.CreatedAt != nil {
		createdAt := user.CreatedAt.AsTime()
		res.CreatedAt = &createdAt
	}

	if user.UpdatedAt != nil {
		updatedAt := user.UpdatedAt.AsTime()
		res.UpdatedAt = &updatedAt
	}

	return &res
}
