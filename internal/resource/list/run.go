// Package list contains the resource listing logic
package list

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gomatic/modern-go-application/internal/resource"
)

// Resource represents a single resource in the list
type Resource struct {
	ID        resource.ID        `json:"id"`
	Name      resource.Name      `json:"name"`
	Status    resource.Status    `json:"status"`
	Tags      []resource.Tag     `json:"tags"`
	CreatedAt resource.CreatedAt `json:"created_at"`
}

// Result holds the result of resource listing
type Result struct {
	Resources       []Resource         `json:"resources"`
	Total           int                `json:"total"`
	Limit           resource.Limit     `json:"limit"`
	Offset          resource.Offset    `json:"offset"`
	IncludePatterns resource.Pattern   `json:"include_patterns,omitempty"`
	ExcludePatterns resource.Pattern   `json:"exclude_patterns,omitempty"`
	FilterStatuses  []resource.Status  `json:"filter_statuses,omitempty"`
	SortBy          resource.SortField `json:"sort_by"`
	Ascending       bool               `json:"ascending"`
}

// MarshalJSON implements json.Marshaler
func (r Result) MarshalJSON() ([]byte, error) {
	type Alias Result
	return json.Marshal((Alias)(r))
}

// Run executes the resource listing logic (noop stub)
func Run(ctx context.Context, logger *slog.Logger, cfg Config) (Result, error) {
	logger.Info("Listing resources",
		"include", cfg.IncludePatterns,
		"exclude", cfg.ExcludePatterns,
		"statuses", cfg.Statuses,
		"limit", cfg.Limit,
		"offset", cfg.Offset,
		"sort_by", cfg.SortBy,
		"ascending", cfg.Ascending,
	)

	// Ensure statuses is not nil
	statuses := cfg.Statuses
	if statuses == nil {
		statuses = []resource.Status{}
	}

	// Generate some mock resources
	resources := []Resource{
		{
			ID:        "res-001",
			Name:      "example-resource-1",
			Status:    resource.StatusActive,
			Tags:      []resource.Tag{"prod", "critical"},
			CreatedAt: resource.CreatedAt(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			ID:        "res-002",
			Name:      "example-resource-2",
			Status:    resource.StatusInactive,
			Tags:      []resource.Tag{"dev", "test"},
			CreatedAt: resource.CreatedAt(time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
		{
			ID:        "res-003",
			Name:      "example-resource-3",
			Status:    resource.StatusActive,
			Tags:      []resource.Tag{"staging"},
			CreatedAt: resource.CreatedAt(time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC)),
		},
	}

	// Filter by statuses if provided
	if len(statuses) > 0 {
		filtered := []Resource{}
		for _, res := range resources {
			for _, status := range statuses {
				if res.Status == status {
					filtered = append(filtered, res)
					break
				}
			}
		}
		resources = filtered
	}

	// Apply limit and offset (convert custom types to int for arithmetic)
	total := len(resources)
	start := int(cfg.Offset)
	end := int(cfg.Offset) + int(cfg.Limit)

	if start >= total {
		resources = []Resource{}
	} else {
		if end > total || int(cfg.Limit) == 0 {
			end = total
		}
		resources = resources[start:end]
	}

	result := Result{
		Resources:       resources,
		Total:           total,
		Limit:           cfg.Limit,
		Offset:          cfg.Offset,
		IncludePatterns: cfg.IncludePatterns,
		ExcludePatterns: cfg.ExcludePatterns,
		FilterStatuses:  statuses,
		SortBy:          cfg.SortBy,
		Ascending:       cfg.Ascending,
	}

	logger.Info("Resource listing complete", "count", len(resources), "total", total)
	return result, nil
}
