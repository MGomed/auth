package model

import (
	"time"

	"github.com/MGomed/auth/consts"
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
	ID   int64
	Name *string
	Role *string
}
