package list

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/gomatic/modern-go-application/internal/resource"
)

// Config holds configuration for listing resources
type Config struct {
	IncludePatterns resource.Pattern   // Patterns to include
	ExcludePatterns resource.Pattern   // Patterns to exclude
	Statuses        []resource.Status  // Filter by statuses
	Limit           resource.Limit     // Maximum number of results
	Offset          resource.Offset    // Offset for pagination
	SortBy          resource.SortField // Sort field
	Ascending       bool               // Sort direction
	Output          app.FilePath       // Output file path
	Logging         log.Config
}

func (c Config) OutputFilePath() app.FilePath { return c.Output }
func (c Config) LoggerConfig() log.Config     { return c.Logging }
