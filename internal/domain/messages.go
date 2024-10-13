package domain

import (
	"time"

	api "github.com/MGomed/auth/pkg/user_api"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// RoleToName maps role with it's name
var RoleToName = map[int32]string{
	0: "UNKNOWN",
	1: "USER",
	2: "ADMIN",
}

// NameToRole maps name with it's role
var NameToRole = map[string]int32{
	"UNKNOWN": 0,
	"USER":    1,
	"ADMIN":   2,
}

type UserInfo struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateRequest domain version of api.CreateRequest
type CreateRequest struct {
	Info *UserInfo
}

// CreateReqFromAPIToDomain converts api.CreateRequest to domain.CreateRequest
func CreateReqFromAPIToDomain(req *api.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Info: &UserInfo{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Role:     RoleToName[int32(req.Role)],
		},
	}
}

// CreateResponse domain version of api.CreateResponse
type CreateResponse struct {
	ID int64
}

// CreateRespToAPIFromDomain converts domain.CreateResponse to api.CreateResponse
func CreateRespToAPIFromDomain(resp *CreateResponse) *api.CreateResponse {
	return &api.CreateResponse{
		Id: resp.ID,
	}
}

// GetRequest domain version of api.GetRequest
type GetRequest struct {
	ID int64
}

// GetReqFromAPIToDomain converts api.GetRequest to domain.GetRequest
func GetReqFromAPIToDomain(req *api.GetRequest) *GetRequest {
	return &GetRequest{
		ID: req.Id,
	}
}

// GetResponse domain version of api.GetResponse
type GetResponse struct {
	Info *UserInfo
}

// GetRespToAPIFromDomain converts domain.GetResponse to api.GetResponse
func GetRespToAPIFromDomain(resp *GetResponse) *api.GetResponse {
	return &api.GetResponse{
		Id:        resp.Info.ID,
		Name:      resp.Info.Name,
		Email:     resp.Info.Email,
		Role:      api.Role(NameToRole[resp.Info.Role]),
		CreatedAt: timestamppb.New(resp.Info.CreatedAt),
		UpdatedAt: timestamppb.New(resp.Info.UpdatedAt),
	}
}

// UpdateRequest domain version of api.UpdateRequest
type UpdateRequest struct {
	Info *UserInfo
}

// UpdateReqFromAPIToDomain converts api.UpdateRequest to domain.UpdateRequest
func UpdateReqFromAPIToDomain(req *api.UpdateRequest) *UpdateRequest {
	return &UpdateRequest{
		Info: &UserInfo{
			ID:   req.Id,
			Name: req.Name.GetValue(),
			Role: RoleToName[int32(req.Role)],
		},
	}
}

// DeleteRequest domain version of api.DeleteRequest
type DeleteRequest struct {
	ID int64
}

// DeleteReqFromAPIToDomain converts api.DeleteRequest to domain.DeleteRequest
func DeleteReqFromAPIToDomain(req *api.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		ID: req.Id,
	}
}
