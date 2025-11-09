package stop

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/gomatic/modern-go-application/internal/service"
)

// Config holds configuration for stopping a service
type Config struct {
	ServiceName service.Name    // Service name
	Force       bool            // Force stop
	Timeout     service.Timeout // Timeout in seconds
	PIDs        []service.PID   // Specific PIDs to stop
	Signal      service.Signal  // Signal to send (SIGTERM, SIGKILL, etc.)
	Output      app.FilePath    // Output file path
	Logging     log.Config
}

func (c Config) OutputFilePath() app.FilePath { return c.Output }
func (c Config) LoggerConfig() log.Config     { return c.Logging }
