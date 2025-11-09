// Package create implements the create command
package create

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/resource/create"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "create"
	usage       = "Create a new resource"
	argsUsage   = "[options]"
	description = `Create a new resource with the specified configuration.

This command demonstrates various configuration types:
  - Strings: name, description
  - Booleans: enabled, dry-run, force
  - String slices: tags

Examples:
  # Create a simple resource
  modern-go-application resource create --name my-resource

  # Create with all options
  modern-go-application resource create \
    --name my-resource \
    --description "My test resource" \
    --tags prod,critical \
    --enabled \
    --dry-run

  # Using environment variables
  MODERN_GO_APP_RESOURCE_CREATE_NAME=my-resource \
  MODERN_GO_APP_RESOURCE_CREATE_ENABLED=true \
  modern-go-application resource create
`
)

// Flag names
const (
	flagName        = "name"
	flagDescription = "description"
	flagTags        = "tags"
	flagEnabled     = "enabled"
	flagDryRun      = "dry-run"
	flagForce       = "force"
)

// Package-level config populated by urfave/cli via Destination
var cfg create.Config

var runAction = create.Run

// Command returns the CLI command for creating resources
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Flags:       flags(prefix),
		Action:      app.Default(cfg, runAction, app.StringSliceConverter(flagTags, &cfg.Tags)),
	}
}

// flags defines all command flags
func flags(prefix app.AppEnvPrefix) []cli.Flag {
	envPrefix := string(prefix) + "RESOURCE_CREATE_"

	baseFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        flagName,
			Aliases:     []string{"n"},
			Usage:       "Resource name",
			EnvVars:     []string{envPrefix + "NAME"},
			Destination: (*string)(&cfg.Name), // Safe: Name is string underneath
		},
		&cli.StringFlag{
			Name:        flagDescription,
			Aliases:     []string{"desc"},
			Usage:       "Resource description",
			EnvVars:     []string{envPrefix + "DESCRIPTION"},
			Value:       "",
			Destination: (*string)(&cfg.Description), // Safe: Description is string underneath
		},
		&cli.StringSliceFlag{
			Name:    flagTags,
			Aliases: []string{"t"},
			Usage:   "Resource tags (can be specified multiple times or comma-separated)",
			EnvVars: []string{envPrefix + "TAGS"},
		},
		&cli.BoolFlag{
			Name:        flagEnabled,
			Aliases:     []string{"e"},
			Usage:       "Enable the resource",
			EnvVars:     []string{envPrefix + "ENABLED"},
			Value:       false,
			Destination: &cfg.Enabled,
		},
		&cli.BoolFlag{
			Name:        flagDryRun,
			Usage:       "Perform a dry run without creating the resource",
			EnvVars:     []string{envPrefix + "DRY_RUN"},
			Value:       false,
			Destination: &cfg.DryRun,
		},
		&cli.BoolFlag{
			Name:        flagForce,
			Aliases:     []string{"f"},
			Usage:       "Force creation even if resource exists",
			EnvVars:     []string{envPrefix + "FORCE"},
			Value:       false,
			Destination: &cfg.Force,
		},
	}

	return app.WithOutputFlags(prefix, &cfg.Output, baseFlags)
}
