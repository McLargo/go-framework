package bootstrap

import (
	"context"
	"errors"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func RunServer() error {
	bootstrap, err := newBootstrap()
	if err != nil {
		return err
	}

	// start the app
	app := startApp(bootstrap)

	// start the server
	return runServer(app, bootstrap)
}

func runServer(app *fiber.App, bootstrap Bootstrap) error {
	// if debug mode, set config watch and start the server
	if *bootstrap.Config.App.Debug {
		watchConfig(app, bootstrap)
	}

	// setting up graceful shutdown
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer done()

	ggroup, gctx := errgroup.WithContext(ctx)

	// start the server on port in a separate goroutine
	ggroup.Go(func() error {
		if err := app.Listen(bootstrap.Config.App.Port); err != nil && !errors.Is(err, context.Canceled) {
			bootstrap.Logger.Error("Could not listen",
				zap.String("Port", bootstrap.Config.App.Port),
				zap.Error(err),
			)

			return err
		}

		return nil
	})

	// graceful shutdown
	<-gctx.Done()
	bootstrap.Logger.Info("Gracefully shut down servers. Exiting...")

	if err := app.Shutdown(); err != nil {
		bootstrap.Logger.Error("App fiber server forced to shutdown", zap.Error(err))
		return err
	}

	return nil
}

func watchConfig(app *fiber.App, bootstrap Bootstrap) {
	var restartOngoing bool
	// channel to notify when the shutdown is complete
	shutdownComplete := make(chan struct{})

	viper.OnConfigChange(func(e fsnotify.Event) {
		if restartOngoing {
			return
		}

		restartOngoing = true

		bootstrap.Logger.Info("Config changed:", zap.String("file", e.Name))

		err := viper.Unmarshal(bootstrap.Config)
		if err != nil {
			bootstrap.Logger.Error("cannot unmarshall to config struct", zap.Error(err))

			restartOngoing = false

			return
		}

		err = bootstrap.Config.Validate()
		if err != nil {
			bootstrap.Logger.Error("cannot validating config:", zap.Error(err))

			restartOngoing = false

			return
		}

		if *bootstrap.Config.App.Verbose {
			bootstrap.Config.Print(bootstrap.Logger)
		}

		go func() {
			bootstrap.Logger.Info("shutting down server to apply new configuration")

			if err := app.Shutdown(); err != nil {
				bootstrap.Logger.Error("App fiber server forced to shutdown", zap.Error(err))

				restartOngoing = false

				return
			}

			shutdownComplete <- struct{}{}
		}()
		// Wait until shutdown is completed
		<-shutdownComplete
		bootstrap.Logger.Info("starting server with new configuration")

		newApp := startApp(bootstrap)
		if err := runServer(newApp, bootstrap); err != nil {
			bootstrap.Logger.Error("Failed to restart server", zap.Error(err))
		}

		restartOngoing = false
	})
	viper.WatchConfig()
}
