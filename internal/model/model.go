package model

import (
	"time"
)

// RoleNames maps api.Role to its string representation
var RoleNames = map[int32]string{
	0: "UNKNOWN",
	1: "USER",
	2: "ADMIN",
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
