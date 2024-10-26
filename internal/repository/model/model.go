package model

import (
	"database/sql"
	"time"
)

// UserInfo represents service_model.UserInfo with db tags
type UserInfo struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      string       `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
