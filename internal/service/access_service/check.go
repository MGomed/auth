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
		return err
	}

	userClaims, ok := claims.(service_model.UserClaims)
	if !ok {
		return service_errors.ErrInvalidToken
	}

	role, ok := consts.AccessibleMap[route]
	if !ok {
		return service_errors.ErrAccessDenied
	}

	if !slices.Contains(role, userClaims.Role) {
		return service_errors.ErrAccessDenied
	}

	return nil
}
