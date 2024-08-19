package main

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func printErr(msg string) {
	logger.Error(msg, "error", ErrNotFound)
}
