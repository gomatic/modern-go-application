package app

import (
	"encoding/json"
	"log/slog"
	"os"
)

// Output writes the result to stdout or a file
func Output(logger *slog.Logger, filePath FilePath, result json.Marshaler) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// If no output file specified, write to stdout
	if filePath == "" {
		_, err = os.Stdout.Write(data)
		if err != nil {
			return err
		}
		_, err = os.Stdout.WriteString("\n")
		return err
	}

	// Write to file
	logger.Info("Writing output to file", "path", filePath)
	return os.WriteFile(string(filePath), append(data, '\n'), 0o600)
}
