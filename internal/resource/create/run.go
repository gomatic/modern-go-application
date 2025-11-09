// Package create implements resource creation logic
package create

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/gomatic/modern-go-application/internal/resource"
)

// Result holds the result of resource creation
type Result struct {
	Success     bool                 `json:"success"`
	ResourceID  resource.ID          `json:"resource_id"`
	Name        resource.Name        `json:"name"`
	Description resource.Description `json:"description"`
	Tags        []resource.Tag       `json:"tags"`
	Enabled     bool                 `json:"enabled"`
	DryRun      bool                 `json:"dry_run"`
	Force       bool                 `json:"force"`
	Message     resource.Message     `json:"message"`
}

// MarshalJSON implements json.Marshaler
func (r Result) MarshalJSON() ([]byte, error) {
	type Alias Result
	return json.Marshal((Alias)(r))
}

// Run executes the resource creation logic (noop stub)
func Run(ctx context.Context, logger *slog.Logger, cfg Config) (Result, error) {
	logger.Info("Creating resource",
		"name", cfg.Name,
		"description", cfg.Description,
		"tags", cfg.Tags,
		"enabled", cfg.Enabled,
		"dry_run", cfg.DryRun,
		"force", cfg.Force,
	)

	result := Result{
		Success:     true,
		ResourceID:  resource.ID("res-" + strings.ReplaceAll(string(cfg.Name), " ", "-")),
		Name:        cfg.Name,
		Description: cfg.Description,
		Tags:        cfg.Tags,
		Enabled:     cfg.Enabled,
		DryRun:      cfg.DryRun,
		Force:       cfg.Force,
		Message:     "Resource created successfully (noop)",
	}

	if cfg.DryRun {
		result.Message = "Dry run: would have created resource (noop)"
	}

	logger.Info("Resource creation complete", "resource_id", result.ResourceID)
	return result, nil
}
