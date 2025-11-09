// Package list implements the resource list command
package list

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/resource/list"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "list"
	usage       = "List resources"
	argsUsage   = "[options]"
	description = `List resources with optional filtering and pagination.

This command demonstrates various configuration types:
  - Strings: include/exclude patterns, sort field
  - Integers: limit, offset
  - Booleans: ascending sort
  - String slices: filter statuses

Examples:
  # List all resources
  modern-go-application resource list

  # List with filtering
  modern-go-application resource list \
    --include "prod-*" \
    --exclude "*-test" \
    --status active,pending \
    --limit 10

  # List with pagination and sorting
  modern-go-application resource list \
    --limit 20 \
    --offset 40 \
    --sort-by name \
    --ascending

  # Using environment variables
  MODERN_GO_APP_RESOURCE_LIST_LIMIT=10 \
  modern-go-application resource list
`
)

// Flag names
const (
	flagInclude   = "include"
	flagExclude   = "exclude"
	flagStatus    = "status"
	flagLimit     = "limit"
	flagOffset    = "offset"
	flagSortBy    = "sort-by"
	flagAscending = "ascending"
)

// Package-level config populated by urfave/cli via Destination
var cfg list.Config

var runAction = list.Run

// Command returns the CLI command for listing resources
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Flags:       flags(prefix),
		Action:      app.Default(cfg, runAction, app.StringSliceConverter(flagStatus, &cfg.Statuses)),
	}
}

// flags defines all command flags
func flags(prefix app.AppEnvPrefix) []cli.Flag {
	envPrefix := string(prefix) + "RESOURCE_LIST_"

	baseFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        flagInclude,
			Aliases:     []string{"i"},
			Usage:       "Include resources matching pattern",
			EnvVars:     []string{envPrefix + "INCLUDE"},
			Destination: (*string)(&cfg.IncludePatterns),
		},
		&cli.StringFlag{
			Name:        flagExclude,
			Aliases:     []string{"x"},
			Usage:       "Exclude resources matching pattern",
			EnvVars:     []string{envPrefix + "EXCLUDE"},
			Destination: (*string)(&cfg.ExcludePatterns),
		},
		&cli.StringSliceFlag{
			Name:    flagStatus,
			Aliases: []string{"s"},
			Usage:   "Filter by status (can be specified multiple times or comma-separated)",
			EnvVars: []string{envPrefix + "STATUSES"},
		},
		&cli.IntFlag{
			Name:        flagLimit,
			Aliases:     []string{"l"},
			Usage:       "Maximum number of results (0 = unlimited)",
			EnvVars:     []string{envPrefix + "LIMIT"},
			Value:       10,
			Destination: (*int)(&cfg.Limit),
		},
		&cli.IntFlag{
			Name:        flagOffset,
			Usage:       "Offset for pagination",
			EnvVars:     []string{envPrefix + "OFFSET"},
			Value:       0,
			Destination: (*int)(&cfg.Offset),
		},
		&cli.StringFlag{
			Name:        flagSortBy,
			Usage:       "Field to sort by",
			EnvVars:     []string{envPrefix + "SORT_BY"},
			Value:       "name",
			Destination: (*string)(&cfg.SortBy),
		},
		&cli.BoolFlag{
			Name:        flagAscending,
			Aliases:     []string{"asc"},
			Usage:       "Sort in ascending order",
			EnvVars:     []string{envPrefix + "ASCENDING"},
			Value:       true,
			Destination: &cfg.Ascending,
		},
	}

	return app.WithOutputFlags(prefix, &cfg.Output, baseFlags)
}
