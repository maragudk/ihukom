package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/maragudk/snorkel"
	"golang.org/x/sync/errgroup"

	"github.com/maragudk/ihukom/http"
)

func main() {
	logger := snorkel.New(snorkel.Options{})
	if err := start(logger); err != nil {
		logger.Log("Error starting", 1, "error", err)
	}
}

func start(logger *snorkel.Logger) error {
	logger.Log("Starting app", 1)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	s := http.NewServer(http.NewServerOptions{
		Log: logger,
	})

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.Start()
	})

	<-ctx.Done()
	logger.Log("Stopping app", 1)

	eg.Go(func() error {
		return s.Stop()
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	logger.Log("Stopped app", 1)

	return nil
}
