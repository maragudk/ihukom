package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/maragudk/env"
	"github.com/maragudk/snorkel"
	"golang.org/x/sync/errgroup"

	"github.com/maragudk/ihukom/http"
	"github.com/maragudk/ihukom/sql"
)

func main() {
	log := snorkel.New(snorkel.Options{})
	if err := start(log); err != nil {
		log.Event("Error starting", 1, "error", err)
	}
}

func start(log *snorkel.Logger) error {
	log.Event("Starting app", 1)

	_ = env.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	db := sql.NewDatabase(sql.NewDatabaseOptions{
		Log:  log,
		Path: env.GetStringOrDefault("DATABASE_PATH", "app.db"),
	})

	if err := db.Connect(); err != nil {
		return err
	}

	s := http.NewServer(http.NewServerOptions{
		DB:  db,
		Log: log,
	})

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.Start()
	})

	<-ctx.Done()
	log.Event("Stopping app", 1)

	eg.Go(func() error {
		return s.Stop()
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	log.Event("Stopped app", 1)

	return nil
}
