// Package service provides types for service configuration.
package service

// Name represents a service name.
type Name string

// Environment represents a deployment environment (dev, staging, prod).
type Environment string

// Common environment constants.
const (
	EnvironmentDev     Environment = "dev"
	EnvironmentStaging Environment = "staging"
	EnvironmentProd    Environment = "prod"
)

// Host represents a server hostname or IP address.
type Host string

// Port represents a server port number.
type Port int

// DatabaseName represents a database name.
type DatabaseName string

// Username represents a username.
type Username string

// Password represents a password.
type Password string

// SSLMode represents a PostgreSQL SSL connection mode.
type SSLMode string

// Common PostgreSQL SSL mode constants.
const (
	SSLModeDisable    SSLMode = "disable"
	SSLModeAllow      SSLMode = "allow"
	SSLModePrefer     SSLMode = "prefer"
	SSLModeRequire    SSLMode = "require"
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)

// Timeout represents a timeout duration in seconds.
type Timeout int

// WorkerCount represents the number of worker processes.
type WorkerCount int

// PID represents a process ID.
type PID int

// Signal represents an OS signal (SIGTERM, SIGKILL, etc.).
type Signal string

// Common signal constants.
const (
	SignalTERM Signal = "SIGTERM"
	SignalKILL Signal = "SIGKILL"
	SignalINT  Signal = "SIGINT"
	SignalHUP  Signal = "SIGHUP"
)
