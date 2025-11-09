// Package resource provides the resource management CLI command.
package resource

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/commands/resource/create"
	"github.com/gomatic/modern-go-application/internal/app/commands/resource/list"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "resource"
	usage       = "Manage resources"
	argsUsage   = "[command]"
	description = `Manage application resources.

This command provides subcommands for creating and listing resources.

Examples:
  # Create a new resource
  modern-go-application resource create --name my-resource --enabled

  # List resources
  modern-go-application resource list --limit 10
`
)

// Command returns the CLI command for resource management (parent command)
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Subcommands: []*cli.Command{
			create.Command(prefix),
			list.Command(prefix),
		},
	}
}
