// Package start contains the service start logic
package start

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/gomatic/modern-go-application/internal/service"
)

// Result holds the result of service start
type Result struct {
	Success     bool                `json:"success"`
	ServiceName service.Name        `json:"service_name"`
	Environment service.Environment `json:"environment"`
	PID         service.PID         `json:"pid"`
	Database    DatabaseConfig      `json:"database"`
	Server      ServerConfig        `json:"server"`
	Workers     service.WorkerCount `json:"workers"`
	EnableCache bool                `json:"enable_cache"`
	Debug       bool                `json:"debug"`
	Message     string              `json:"message"`
}

// MarshalJSON implements json.Marshaler
func (r Result) MarshalJSON() ([]byte, error) {
	type Alias Result
	return json.Marshal((Alias)(r))
}

// Run executes the service start logic (noop stub)
func Run(ctx context.Context, logger *slog.Logger, cfg Config) (Result, error) {
	logger.Info("Starting service",
		"service_name", cfg.ServiceName,
		"environment", cfg.Environment,
		"workers", cfg.Workers,
		"enable_cache", cfg.EnableCache,
		"debug", cfg.Debug,
	)

	logger.Debug("Database configuration",
		"host", cfg.Database.Host,
		"port", cfg.Database.Port,
		"name", cfg.Database.Name,
		"user", cfg.Database.User,
	)

	logger.Debug("Server configuration",
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
		"read_timeout", cfg.Server.ReadTimeout,
		"write_timeout", cfg.Server.WriteTimeout,
	)

	result := Result{
		Success:     true,
		ServiceName: cfg.ServiceName,
		Environment: cfg.Environment,
		PID:         12345, // Mock PID
		Database:    cfg.Database,
		Server:      cfg.Server,
		Workers:     cfg.Workers,
		EnableCache: cfg.EnableCache,
		Debug:       cfg.Debug,
		Message:     "Service started successfully (noop)",
	}

	logger.Info("Service start complete", "service_name", result.ServiceName, "pid", result.PID)
	return result, nil
}
