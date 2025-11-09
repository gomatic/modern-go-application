package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sort"

	"github.com/gomatic/modern-go-application/internal/app/commands/resource"
	"github.com/gomatic/modern-go-application/internal/app/commands/service"
	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/urfave/cli/v2"
)

const (
	appName      = "mga"
	appEnvName   = "MGA"
	appUsage     = "A modern Go application demonstrating best practices"
	appEnvPrefix = appEnvName + "_"
)

var appVersion string

var loggerConfig log.Config

func main() { run() }

var (
	appCreator    = createApp
	loggerCreator = log.GetLogger
)

func run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	c := appCreator(loggerCreator)

	if err := c.RunContext(ctx, os.Args); err != nil {
		slog.Error("Application error", "error", err)
		cancel()
		os.Exit(1)
	}
}

func createApp(getLogger log.GetLoggerFunc) *cli.App {
	c := &cli.App{
		Name:    appName,
		Usage:   appUsage,
		Version: appVersion,
		Commands: []*cli.Command{
			resource.Command(appEnvPrefix),
			service.Command(appEnvPrefix),
		},
		Before: func(c *cli.Context) error {
			c.App.Metadata[log.LoggerMetadataKey] = getLogger(c, loggerConfig)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				EnvVars:     []string{appEnvPrefix + "LOG_LEVEL"},
				Value:       "info",
				Usage:       "Set the logging level (debug, info, warn, error)",
				Destination: (*string)(&loggerConfig.Level),
			},
			&cli.StringFlag{
				Name:        "log-format",
				EnvVars:     []string{appEnvPrefix + "LOG_FORMAT"},
				Value:       "text",
				Usage:       "Set the log output format (text, json)",
				Destination: (*string)(&loggerConfig.Format),
			},
		},
	}

	// Sort commands alphabetically
	sort.Sort(cli.CommandsByName(c.Commands))

	return c
}
