// Package start implements the "service start" command.
package start

import (
	"github.com/gomatic/modern-go-application/internal/app"
	"github.com/gomatic/modern-go-application/internal/service/start"
	"github.com/urfave/cli/v2"
)

// Command metadata
const (
	Name        = "start"
	usage       = "Start a service"
	argsUsage   = "[options]"
	description = `Start a service with the specified configuration.

This command demonstrates nested configuration structures:
  - Database configuration (host, port, name, user, password)
  - Server configuration (host, port, timeouts)
  - Other settings (workers, cache, debug)

Examples:
  # Start with default configuration
  modern-go-application service start --service-name my-service

  # Start with custom database configuration
  modern-go-application service start \
    --service-name my-service \
    --environment prod \
    --db-host postgres.example.com \
    --db-port 5432 \
    --db-name mydb \
    --db-user admin

  # Start with custom server configuration
  modern-go-application service start \
    --service-name my-service \
    --server-host 0.0.0.0 \
    --server-port 8080 \
    --workers 4 \
    --enable-cache \
    --debug

  # Using environment variables
  MGA_START_SERVICE_NAME=my-service \
  MGA_START_DB_HOST=localhost \
  MGA_START_SERVER_PORT=8080 \
  modern-go-application service start
`
)

// Flag names
const (
	flagServiceName        = "service-name"
	flagEnvironment        = "environment"
	flagDBHost             = "db-host"
	flagDBPort             = "db-port"
	flagDBName             = "db-name"
	flagDBUser             = "db-user"
	flagDBPassword         = "db-password"
	flagDBSSLMode          = "db-sslmode"
	flagServerHost         = "server-host"
	flagServerPort         = "server-port"
	flagServerReadTimeout  = "server-read-timeout"
	flagServerWriteTimeout = "server-write-timeout"
	flagWorkers            = "workers"
	flagEnableCache        = "enable-cache"
	flagDebug              = "debug"
)

// Package-level config populated by urfave/cli via Destination
var cfg start.Config

var runAction = start.Run

// Command returns the CLI command for starting services
func Command(prefix app.AppEnvPrefix) *cli.Command {
	return &cli.Command{
		Name:        Name,
		Usage:       usage,
		ArgsUsage:   argsUsage,
		Description: description,
		Flags:       flags(prefix),
		Action:      app.Default(cfg, runAction),
	}
}

// flags defines all command flags
func flags(prefix app.AppEnvPrefix) []cli.Flag {
	envPrefix := string(prefix) + "SERVICE_START_"

	baseFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        flagServiceName,
			Aliases:     []string{"n"},
			Usage:       "Service name",
			EnvVars:     []string{envPrefix + "SERVICE_NAME"},
			Destination: (*string)(&cfg.ServiceName),
		},
		&cli.StringFlag{
			Name:        flagEnvironment,
			Aliases:     []string{"env"},
			Usage:       "Environment (dev, staging, prod)",
			EnvVars:     []string{envPrefix + "ENVIRONMENT"},
			Value:       "dev",
			Destination: (*string)(&cfg.Environment),
		},

		// Database configuration (uses official PostgreSQL env vars)
		&cli.StringFlag{
			Name:        flagDBHost,
			Usage:       "Database host",
			EnvVars:     []string{"PGHOST"},
			Value:       "localhost",
			Destination: (*string)(&cfg.Database.Host),
		},
		&cli.IntFlag{
			Name:        flagDBPort,
			Usage:       "Database port",
			EnvVars:     []string{"PGPORT"},
			Value:       5432,
			Destination: (*int)(&cfg.Database.Port),
		},
		&cli.StringFlag{
			Name:        flagDBName,
			Usage:       "Database name",
			EnvVars:     []string{"PGDATABASE"},
			Value:       "postgres",
			Destination: (*string)(&cfg.Database.Name),
		},
		&cli.StringFlag{
			Name:        flagDBUser,
			Usage:       "Database user",
			EnvVars:     []string{"PGUSER"},
			Value:       "postgres",
			Destination: (*string)(&cfg.Database.User),
		},
		&cli.StringFlag{
			Name:        flagDBPassword,
			Usage:       "Database password",
			EnvVars:     []string{"PGPASSWORD"},
			Destination: (*string)(&cfg.Database.Password),
		},
		&cli.StringFlag{
			Name:        flagDBSSLMode,
			Usage:       "SSL mode (disable, require, verify-ca, verify-full)",
			EnvVars:     []string{"PGSSLMODE"},
			Value:       "disable", // Local dev default - override in production
			Destination: (*string)(&cfg.Database.SSLMode),
		},

		// Server configuration
		&cli.StringFlag{
			Name:        flagServerHost,
			Usage:       "Server host",
			EnvVars:     []string{envPrefix + "SERVER_HOST"},
			Value:       "localhost",
			Destination: (*string)(&cfg.Server.Host),
		},
		&cli.IntFlag{
			Name:        flagServerPort,
			Usage:       "Server port",
			EnvVars:     []string{envPrefix + "SERVER_PORT"},
			Value:       8080,
			Destination: (*int)(&cfg.Server.Port),
		},
		&cli.IntFlag{
			Name:        flagServerReadTimeout,
			Usage:       "Server read timeout (seconds)",
			EnvVars:     []string{envPrefix + "SERVER_READ_TIMEOUT"},
			Value:       30,
			Destination: (*int)(&cfg.Server.ReadTimeout),
		},
		&cli.IntFlag{
			Name:        flagServerWriteTimeout,
			Usage:       "Server write timeout (seconds)",
			EnvVars:     []string{envPrefix + "SERVER_WRITE_TIMEOUT"},
			Value:       30,
			Destination: (*int)(&cfg.Server.WriteTimeout),
		},

		// Other configuration
		&cli.IntFlag{
			Name:        flagWorkers,
			Aliases:     []string{"w"},
			Usage:       "Number of worker processes",
			EnvVars:     []string{envPrefix + "WORKERS"},
			Value:       2,
			Destination: (*int)(&cfg.Workers),
		},
		&cli.BoolFlag{
			Name:        flagEnableCache,
			Usage:       "Enable caching",
			EnvVars:     []string{envPrefix + "ENABLE_CACHE"},
			Value:       false,
			Destination: &cfg.EnableCache,
		},
		&cli.BoolFlag{
			Name:        flagDebug,
			Aliases:     []string{"d"},
			Usage:       "Enable debug mode",
			EnvVars:     []string{envPrefix + "DEBUG"},
			Value:       false,
			Destination: &cfg.Debug,
		},
	}

	return app.WithOutputFlags(prefix, &cfg.Output, baseFlags)
}
