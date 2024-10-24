package config

// APIConfig is grpc config interface
type APIConfig interface {
	Address() string
}

// PgConfig is postgres config interface
type PgConfig interface {
	DSN() string
}

// LoggerConfig is logger config interface
type LoggerConfig interface {
	OutDir() string
}
