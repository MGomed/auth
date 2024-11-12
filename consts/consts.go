package consts

import "time"

// ContextTimeout is timeout for db connecting and server start
const ContextTimeout = 15 * time.Second

// Server env's names
const (
	ServerHostEnv = "SERVER_HOST"
	ServerPortEnv = "SERVER_PORT"
)

// DB env's names
const (
	DBHostEnv     = "DB_HOST"
	DBPortEnv     = "DB_PORT"
	DBNameEnv     = "POSTGRES_DB"
	DBUserEnv     = "POSTGRES_USER"
	DBPasswordEnv = "POSTGRES_PASSWORD" //nolint: gosec
)

// Redis env's names
const (
	RedisHostEnv              = "REDIS_HOST"
	RedisPortEnv              = "REDIS_PORT"
	RedisConnectionTimeoutEnv = "REDIS_CONNECTION_TIMEOUT_SEC"
	RedisMaxIdleEnv           = "REDIS_MAX_IDLE"
	RedisIdleTimeoutEnv       = "REDIS_IDLE_TIMEOUT_SEC"
)

// DB table and columns names
const (
	AuthTable = "auth"

	IDColumn        = "id"
	NameColumn      = "name"
	EmailColumn     = "email"
	PasswordColumn  = "password"
	RoleColumn      = "role"
	CreatedAtColumn = "created_at"
	UpdatedAtColumn = "updated_at"
)

// Role const names
const (
	RoleUnknown = "UNKNOWN"
	RoleUser    = "USER"
	RoleAdmin   = "ADMIN"
)
