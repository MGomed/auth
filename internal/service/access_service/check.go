package accessservice

import (
	"slices"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	service_errors "github.com/MGomed/auth/internal/service/errors"
	token "github.com/MGomed/common/token"
)

func (s *service) Check(route string, accessToken string, secretKey []byte) error {
	claims, err := token.VerifyToken(accessToken, secretKey, &service_model.UserClaims{})
	if err != nil {
		s.log.Error("Failed to verify refresh token: %v", err)

		return err
	}

	userClaims, ok := claims.(service_model.UserClaims)
	if !ok {
		s.log.Error("Wrong claims in token")

		return service_errors.ErrInvalidToken
	}

	role, ok := consts.AccessibleMap[route]
	if !ok {
		s.log.Error("Unknown route for access")

		return service_errors.ErrAccessDenied
	}

	if !slices.Contains(role, userClaims.Role) {
		s.log.Error("Access denied for user with email: %v", userClaims.Email)

		return service_errors.ErrAccessDenied
	}

	return nil
}
