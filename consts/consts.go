package consts

import "time"

// ServiceName is a application name
const ServiceName = "auth_service"

// Security files path
const (
	ServiceCertFilePath    = "certs/service.pem"
	ServiceCertKeyFilePath = "certs/service.key"
)

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

// Prometheus Server env's names
const (
	PrometheusServerHostEnv = "PROMETHEUS_HOST"
	PrometheusServerPortEnv = "PROMETHEUS_PORT"
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

// JWT env's names
const (
	JWTRefreshTokenExpirationTimeMinEnv = "REFRESH_TOKEN_EXPIRATION_MIN" //nolint:gosec
	JWTAccessTokenExpirationTimeMinEnv  = "ACCESS_TOKEN_EXPIRATION_MIN"  //nolint:gosec
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

// AccessibleMap defines access to different route for user roles
var AccessibleMap = map[string][]string{
	"/chat_service/CreateChat":    {RoleAdmin},
	"/chat_service/ConnectToChat": {RoleAdmin, RoleUser},
	"/chat_service/SendMessage":   {RoleAdmin, RoleUser},
}

// AccessPrefix defines access prefix in grpc context
var AccessPrefix = "Bearer "

// Secret key paths
const (
	RefreshSecretKeyPath = "certs/refresh_secret_key" //nolint:gosec
	AccessSecretKeyPath  = "certs/access_secret_key"  //nolint:gosec
)

// Metric's consts
const (
	MetricsNamespace = "auth_namespace"
	MetricsAppName   = "auth"
)

// Metric's labels const
const (
	MetricStatusLabel = "status"
	MetricMethodLabel = "method"
)
