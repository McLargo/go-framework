package bootstrap

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mclargo/go-framework/internal/handlers"

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

	// initialize a new Fiber app
	app := fiber.New()

	// set endpoint -> healthz
	app.Get("/healthz", handlers.Healthz)

	if *bootstrap.Config.App.Verbose {
		for _, route := range app.GetRoutes() {
			bootstrap.Logger.Debug("Endpoint registered",
				zap.String("Route", route.Path),
				zap.String("Method", route.Method),
			)
		}
	}

	return startServer(app, bootstrap)
}

func startServer(app *fiber.App, bootstrap Bootstrap) error {
	// setting up graceful shutdown
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer done()

	ggroup, gctx := errgroup.WithContext(ctx)

	// Channel to ensure graceful shutdown completes before restarting
	shutdownComplete := make(chan struct{})
	defer close(shutdownComplete)

	var restartOngoing bool
	// watch for config changes
	if *bootstrap.Config.App.Debug {
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
			restartOngoing = false

			if err := startServer(app, bootstrap); err != nil {
				bootstrap.Logger.Error("Failed to restart server", zap.Error(err))
			}

		})

		viper.WatchConfig()
	}

	// start the server on port in a separate goroutine
	ggroup.Go(func() error {
		if err := app.Listen(bootstrap.Config.App.Port); err != nil && err != context.Canceled {
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
