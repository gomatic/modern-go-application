// Package app provides the application framework including action handlers,
// flag helpers, and output formatting.
package app

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/gomatic/modern-go-application/internal/app/log"
	"github.com/gomatic/modern-go-application/internal/slice"
	"github.com/urfave/cli/v2"
)

// Configurable interface for command configurations
type Configurable interface {
	HasLoggerConfig
	HasOutput
}

// HasLoggerConfig interface for configs that have logger configuration
type HasLoggerConfig interface {
	LoggerConfig() log.Config
}

// HasOutput interface for configs that have output configuration
type HasOutput interface {
	OutputFilePath() FilePath
}

// Runner is a generic function type for command runners
type Runner[CONFIG Configurable, RESULT json.Marshaler] func(context.Context, *slog.Logger, CONFIG) (RESULT, error)

var getLogger = log.GetLogger

// Action is a generic action handler that executes a runner and outputs the result
func Action[C Configurable, R json.Marshaler](c *cli.Context, cfg C, runner Runner[C, R]) error {
	logger := getLogger(c, cfg.LoggerConfig())

	result, err := runner(c.Context, logger, cfg)
	if err != nil {
		return err
	}

	return Output(logger, cfg.OutputFilePath(), result)
}

func Default[C Configurable, R json.Marshaler](cfg C, runner Runner[C, R], options ...any) cli.ActionFunc {
	return func(c *cli.Context) error {
		for _, opt := range options {
			switch o := opt.(type) {
			case converter:
				o.Convert(c)
			default:
				slog.Warn("Unknown option type", "type", o)
			}
		}
		return Action(c, cfg, runner)
	}
}

type converter interface {
	Convert(c *cli.Context)
}

// sliceConverter handles conversion from CLI slice types to custom domain slice types.
type sliceConverter[F, T any] struct {
	flagName string // Name of the CLI flag
	slicer   slicerFunc[F]
	dest     *[]T                    // Pointer to destination slice in config
	convert  slice.ConvertFunc[F, T] // Conversion function from CLI type to custom type
}

func (s sliceConverter[F, T]) Convert(c *cli.Context) {
	*s.dest = slice.Convert[F, T](s.slicer(c, s.flagName), s.convert)
}

// StringSliceConverter creates a converter for string slices with automatic type conversion.
// The conversion defaults to casting: I(s), which works for all string-based custom types.
func StringSliceConverter[T ~string](flagName string, dest *[]T) sliceConverter[string, T] {
	return sliceConverter[string, T]{
		flagName: flagName,
		slicer:   stringSlice,
		dest:     dest,
		convert:  slice.StringConverter[string, T],
	}
}

// IntSliceConverter creates a converter for int slices with automatic type conversion.
// The conversion defaults to casting: I(i), which works for all int-based custom types.
func IntSliceConverter[T ~int](flagName string, dest *[]T) sliceConverter[int, T] {
	return sliceConverter[int, T]{
		flagName: flagName,
		slicer:   intSlice,
		dest:     dest,
		convert:  slice.IntConverter[int, T],
	}
}

type slicerFunc[T any] func(*cli.Context, string) []T

func intSlice(c *cli.Context, flagName string) []int       { return c.IntSlice(flagName) }
func stringSlice(c *cli.Context, flagName string) []string { return c.StringSlice(flagName) }
