package consts

import "time"

// Timeouts uses in project
const (
	ContextTimeout    = 30 * time.Second
	ReadHeaderTimeout = 5 * time.Second
)

// GRPC Server env's names
const (
	GRPCServerHostEnv = "GRPC_HOST"
	GRPCServerPortEnv = "GRPC_PORT"
)

// HTTP Server env's names
const (
	HTTPServerHostEnv = "HTTP_HOST"
	HTTPServerPortEnv = "HTTP_PORT"
)

// Swagger Server env's names
const (
	SwaggerServerHostEnv = "SWAGGER_HOST"
	SwaggerServerPortEnv = "SWAGGER_PORT"
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
	RedisHostEnv                 = "REDIS_HOST"
	RedisPortEnv                 = "REDIS_PORT"
	RedisConnectionTimeoutSecEnv = "REDIS_CONNECTION_TIMEOUT_SEC"
	RedisMaxIdleEnv              = "REDIS_MAX_IDLE"
	RedisIdleTimeoutSecEnv       = "REDIS_IDLE_TIMEOUT_SEC"
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

// SwaggerPath declaires swagger json file path/url
const SwaggerPath = "/api.swagger.json"

// Kafka topic names
const (
	CreateTopic string = "create_topic"
	DeleteTopic string = "delete_topic"
)

// Kafka env's names
const (
	KafkaBrokersEnvName = "KAFKA_BROKERS"
	KafkaGroupIDEnvName = "KAFKA_GROUP_ID"
)
