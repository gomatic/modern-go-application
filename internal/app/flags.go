package app

import (
	"github.com/urfave/cli/v2"
)

// AppEnvPrefix is a type for application environment variable prefixes
type AppEnvPrefix string

// WithOutputFlags appends output flags to the provided flag list
func WithOutputFlags(prefix AppEnvPrefix, output *FilePath, flags []cli.Flag) []cli.Flag {
	return append(flags, OutputFlags(prefix, output)...)
}

// OutputFlags returns standard output flags
func OutputFlags(prefix AppEnvPrefix, output *FilePath) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output file path (default: stdout)",
			EnvVars:     []string{string(prefix) + "OUTPUT"},
			Destination: (*string)(output), // Safe: FilePath is string underneath
		},
	}
}

// WithCommonFlags appends common flags (debug, verbose, etc.) to the provided flag list
func WithCommonFlags(prefix AppEnvPrefix, flags []cli.Flag, debugDest *bool, verboseDest *bool) []cli.Flag {
	commonFlags := []cli.Flag{}

	if debugDest != nil {
		commonFlags = append(commonFlags, &cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Usage:       "Enable debug mode",
			EnvVars:     []string{string(prefix) + "DEBUG"},
			Destination: debugDest,
		})
	}

	if verboseDest != nil {
		commonFlags = append(commonFlags, &cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			Usage:       "Enable verbose output",
			EnvVars:     []string{string(prefix) + "VERBOSE"},
			Destination: verboseDest,
		})
	}

	return append(flags, commonFlags...)
}

// WithFilterFlags appends filter flags to the provided flag list
func WithFilterFlags(prefix AppEnvPrefix, includePatterns *string, excludePatterns *string, flags []cli.Flag) []cli.Flag {
	filterFlags := []cli.Flag{}

	if includePatterns != nil {
		filterFlags = append(filterFlags, &cli.StringFlag{
			Name:        "include",
			Usage:       "Comma-separated patterns to include",
			EnvVars:     []string{string(prefix) + "INCLUDE"},
			Destination: includePatterns,
		})
	}

	if excludePatterns != nil {
		filterFlags = append(filterFlags, &cli.StringFlag{
			Name:        "exclude",
			Usage:       "Comma-separated patterns to exclude",
			EnvVars:     []string{string(prefix) + "EXCLUDE"},
			Destination: excludePatterns,
		})
	}

	return append(flags, filterFlags...)
}
