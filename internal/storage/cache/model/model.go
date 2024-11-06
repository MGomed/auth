package model

import "time"

// UserInfo represents service_model.UserInfo with redis tags
type UserInfo struct {
	Name      string    `redis:"name"`
	Email     string    `redis:"email"`
	Password  string    `redis:"password"`
	Role      string    `redis:"role"`
	CreatedAt time.Time `redis:"created_at"`
	UpdatedAt time.Time `redis:"updated_at"`
}
