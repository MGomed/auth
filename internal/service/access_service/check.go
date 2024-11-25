package accessservice

import (
	"errors"
	"slices"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	token "github.com/MGomed/auth/pkg/token"
)

func (s *service) Check(route string, accessToken string, secretKey []byte) error {
	claims, err := token.VerifyToken(accessToken, secretKey, &service_model.UserClaims{})
	if err != nil {
		return err
	}

	userClaims, ok := claims.(service_model.UserClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	role, ok := consts.AccessibleMap[route]
	if !ok {
		return errors.New("access denied")
	}

	if !slices.Contains(role, userClaims.Role) {
		return errors.New("access denied")
	}

	return nil
}
