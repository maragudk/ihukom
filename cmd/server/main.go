package main

import (
	"github.com/maragudk/snorkel"
)

func main() {
	logger := snorkel.New(snorkel.Options{})
	if err := start(logger); err != nil {
		logger.Log("Error starting", 1, "error", err)
	}
}

func start(logger *snorkel.Logger) error {
	logger.Log("Starting", 1)
	return nil
}
