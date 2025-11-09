// Package service provides the CLI command for service management
package service

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/app/commands/service/start"
	"github.com/gomatic/modern-go-application/internal/app/commands/service/stop"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "service"
	usage       = "Manage services"
	argsUsage   = "[command]"
	description = `Manage application services.

This command provides subcommands for starting and stopping services.

Examples:
  # Start a service
  modern-go-application service start --service-name my-service

  # Stop a service
  modern-go-application service stop --service-name my-service
`
)

// Command returns the CLI command for service management (parent command)
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Subcommands: []*cli.Command{
			start.Command(prefix),
			stop.Command(prefix),
		},
	}
}
