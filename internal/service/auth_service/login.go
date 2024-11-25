package authservice

import (
	"context"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"

	service_model "github.com/MGomed/auth/internal/model"
	token "github.com/MGomed/auth/pkg/token"
	jwt "github.com/dgrijalva/jwt-go"
)

// Login generates refreshToken for authorized user
func (s *service) Login(
	ctx context.Context,
	email, password string,
	secretKey []byte,
	duration time.Duration,
) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return "", err
	}

	refreshToken, err := token.GenerateToken(service_model.UserClaims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}, secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
