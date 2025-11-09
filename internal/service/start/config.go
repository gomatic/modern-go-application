package start

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/gomatic/modern-go-application/internal/service"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     service.Host
	Port     service.Port
	Name     service.DatabaseName
	User     service.Username
	Password service.Password
	SSLMode  service.SSLMode
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Host         service.Host
	Port         service.Port
	ReadTimeout  service.Timeout
	WriteTimeout service.Timeout
}

// Config holds configuration for starting a service (nested config example)
type Config struct {
	ServiceName service.Name        // Service name
	Environment service.Environment // Environment (dev, staging, prod)
	Database    DatabaseConfig      // Nested database config
	Server      ServerConfig        // Nested server config
	Workers     service.WorkerCount // Number of workers
	EnableCache bool                // Enable caching
	Debug       bool                // Debug mode
	Output      app.FilePath        // Output file path
	Logging     log.Config
}

func (c Config) OutputFilePath() app.FilePath { return c.Output }
func (c Config) LoggerConfig() log.Config     { return c.Logging }
