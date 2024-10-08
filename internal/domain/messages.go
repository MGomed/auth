package domain

import (
	"time"

	api "github.com/MGomed/auth/pkg/user_api"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Role domain version of api.Role
type Role int32

// CreateRequest domain version of api.CreateRequest
type CreateRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

// CreateReqFromAPIToDomain converts api.CreateRequest to domain.CreateRequest
func CreateReqFromAPIToDomain(req *api.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Role:            Role(req.Role),
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
	ID        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetRespToAPIFromDomain converts domain.GetResponse to api.GetResponse
func GetRespToAPIFromDomain(resp *GetResponse) *api.GetResponse {
	return &api.GetResponse{
		Id:        resp.ID,
		Name:      resp.Name,
		Email:     resp.Email,
		Role:      api.Role(resp.Role),
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

// UpdateRequest domain version of api.UpdateRequest
type UpdateRequest struct {
	ID    int64
	Name  string
	Email string
	Role  Role
}

// UpdateReqFromAPIToDomain converts api.UpdateRequest to domain.UpdateRequest
func UpdateReqFromAPIToDomain(req *api.UpdateRequest) *UpdateRequest {
	return &UpdateRequest{
		ID:   req.Id,
		Name: req.Name.GetValue(),
		Role: Role(req.Role),
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
