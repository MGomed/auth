package accessapiimpl

import (
	"context"
	"errors"
	"strings"

	empty "github.com/golang/protobuf/ptypes/empty"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	consts "github.com/MGomed/auth/consts"
	access_api "github.com/MGomed/auth/pkg/access_api"
)

// Check checks that the user have right access to route
func (api *AccessAPI) Check(ctx context.Context, req *access_api.CheckRequest) (*empty.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], consts.AccessPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], consts.AccessPrefix)

	if err := api.service.Check(req.EndpointAddress, accessToken, api.accessSecretKey); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
