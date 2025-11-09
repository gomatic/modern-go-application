// Package stop implements the stop command
package stop

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/service/stop"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "stop"
	usage       = "Stop a service"
	argsUsage   = "[options]"
	description = `Stop a running service with configurable options.

This command demonstrates various configuration types:
  - Strings: service name, signal
  - Integers: timeout
  - Integer slices: PIDs
  - Booleans: force

Examples:
  # Stop a service
  modern-go-application service stop --service-name myservice

  # Stop with force and custom timeout
  modern-go-application service stop \
    --service-name myservice \
    --force \
    --timeout 60

  # Stop specific PIDs
  modern-go-application service stop \
    --service-name myservice \
    --pid 1234 --pid 5678 \
    --signal SIGKILL

  # Using environment variables
  MODERN_GO_APP_SERVICE_STOP_SERVICE_NAME=myservice \
  MODERN_GO_APP_SERVICE_STOP_FORCE=true \
  modern-go-application service stop
`
)

// Flag names
const (
	flagServiceName = "service-name"
	flagForce       = "force"
	flagTimeout     = "timeout"
	flagPID         = "pid"
	flagSignal      = "signal"
)

// Package-level config populated by urfave/cli via Destination
var cfg stop.Config

var runAction = stop.Run

// Command returns the CLI command for stopping services
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Flags:       flags(prefix),
		Action:      app.Default(cfg, runAction, app.IntSliceConverter(flagPID, &cfg.PIDs)),
	}
}

// flags defines all command flags
func flags(prefix app.AppEnvPrefix) []cli.Flag {
	envPrefix := string(prefix) + "SERVICE_STOP_"

	baseFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        flagServiceName,
			Aliases:     []string{"n"},
			Usage:       "Service name",
			EnvVars:     []string{envPrefix + "SERVICE_NAME"},
			Destination: (*string)(&cfg.ServiceName),
		},
		&cli.BoolFlag{
			Name:        flagForce,
			Aliases:     []string{"f"},
			Usage:       "Force stop the service",
			EnvVars:     []string{envPrefix + "FORCE"},
			Value:       false,
			Destination: &cfg.Force,
		},
		&cli.IntFlag{
			Name:        flagTimeout,
			Aliases:     []string{"t"},
			Usage:       "Timeout in seconds",
			EnvVars:     []string{envPrefix + "TIMEOUT"},
			Value:       30,
			Destination: (*int)(&cfg.Timeout),
		},
		&cli.IntSliceFlag{
			Name:    string(flagPID),
			Aliases: []string{"p"},
			Usage:   "Specific PID to stop (can be specified multiple times)",
			EnvVars: []string{envPrefix + "PIDS"},
		},
		&cli.StringFlag{
			Name:        flagSignal,
			Aliases:     []string{"s"},
			Usage:       "Signal to send (SIGTERM, SIGKILL, SIGINT, etc.)",
			EnvVars:     []string{envPrefix + "SIGNAL"},
			Value:       "SIGTERM",
			Destination: (*string)(&cfg.Signal),
		},
	}

	return app.WithOutputFlags(prefix, &cfg.Output, baseFlags)
}
