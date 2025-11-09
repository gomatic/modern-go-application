// Package stop implements the service stop logic
package stop

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/gomatic/modern-go-application/internal/service"
)

// Result holds the result of service stop
type Result struct {
	Success     bool            `json:"success"`
	ServiceName service.Name    `json:"service_name"`
	Force       bool            `json:"force"`
	Timeout     service.Timeout `json:"timeout"`
	PIDs        []service.PID   `json:"pids"`
	Signal      service.Signal  `json:"signal"`
	Message     string          `json:"message"`
}

// MarshalJSON implements json.Marshaler
func (r Result) MarshalJSON() ([]byte, error) {
	type Alias Result
	return json.Marshal((Alias)(r))
}

// Run executes the service stop logic (noop stub)
func Run(ctx context.Context, logger *slog.Logger, cfg Config) (Result, error) {
	logger.Info("Stopping service",
		"service_name", cfg.ServiceName,
		"force", cfg.Force,
		"timeout", cfg.Timeout,
		"pids", cfg.PIDs,
		"signal", cfg.Signal,
	)

	result := Result{
		Success:     true,
		ServiceName: cfg.ServiceName,
		Force:       cfg.Force,
		Timeout:     cfg.Timeout,
		PIDs:        cfg.PIDs,
		Signal:      cfg.Signal,
		Message:     "Service stopped successfully (noop)",
	}

	if cfg.Force {
		result.Message = "Service force stopped successfully (noop)"
	}

	logger.Info("Service stop complete", "service_name", result.ServiceName)
	return result, nil
}
