package model

import (
	"time"

	"github.com/MGomed/auth/consts"
	"github.com/dgrijalva/jwt-go"
)

// RoleNames maps api.Role to its string representation
var RoleNames = map[int32]string{
	0: consts.RoleUnknown,
	1: consts.RoleUser,
	2: consts.RoleAdmin,
}

// UserCreate represent api.UserCreate object
type UserCreate struct {
	Name     string
	Email    string
	Password []byte
	Role     string
}

// UserInfo represent api.UserInfo object
type UserInfo struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

// UserUpdate represent api.UserUpdate object
type UserUpdate struct {
	Name *string
	Role *string
}

// UserClaims is jwt.Claims for auth service
type UserClaims struct {
	jwt.StandardClaims

	Email string `json:"email"`
	Role  string `json:"role"`
}

// Tokens declare struct with refresh and access tokens
type Tokens struct {
	RefreshToken string
	AccessToken  string
}
