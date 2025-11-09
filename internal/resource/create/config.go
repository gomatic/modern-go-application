package create

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/gomatic/modern-go-application/internal/resource"
)

// Config holds configuration for resource creation
type Config struct {
	Name        resource.Name        // Resource name
	Description resource.Description // Resource description
	Tags        []resource.Tag       // Resource tags
	Enabled     bool                 // Whether resource is enabled (bool example)
	DryRun      bool                 // Dry run mode
	Force       bool                 // Force creation
	Output      app.FilePath         // Output file path
	Logging     log.Config
}

func (c Config) OutputFilePath() app.FilePath { return c.Output }
func (c Config) LoggerConfig() log.Config     { return c.Logging }
