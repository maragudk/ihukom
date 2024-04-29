package main

import (
	"log/slog"
)

func main() {
	if err := start(); err != nil {
		slog.Info("Error starting:", "error", err)
	}
}

func start() error {
	slog.Info("Starting")
	return nil
}
