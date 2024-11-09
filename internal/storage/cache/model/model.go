package model

// UserInfo represents service_model.UserInfo with redis tags
type UserInfo struct {
	Name      string `redis:"name"`
	Email     string `redis:"email"`
	Password  string `redis:"password"`
	Role      string `redis:"role"`
	CreatedAt int64  `redis:"created_at"`
	UpdatedAt int64  `redis:"updated_at"`
}
