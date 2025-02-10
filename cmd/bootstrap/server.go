package bootstrap

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

func RunServer() error {
	cfg, err := conf.InitConfig()
	if err != nil {
		return err
	}

	if *cfg.App.Verbose {
		cfg.Print()
	}

	if *cfg.App.Debug {
		// Watch the config file for changes
		cfg.Watch()
	}

	// TODO: init logger

	// initialize a new Fiber app
	app := fiber.New()

	// set endpoint -> healthz
	app.Get("/healthz", handlers.Healthz)

	// setting up graceful shutdown
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer done()

	ggroup, gctx := errgroup.WithContext(ctx)

	// start the server on port in a separate goroutine
	ggroup.Go(func() error {
		if err := app.Listen(cfg.App.Port); err != nil && err != context.Canceled {
			// TODO: log error
			//logger.Errorf("Could not listen on 3000: %v", err)
			return err
		}
		return nil
	})

	// graceful shutdown
	<-gctx.Done()
	// TODO: log info
	fmt.Println("Gracefully shut down servers. Exiting...")
	if err := app.Shutdown(); err != nil {
		// TODO: log error
		//logger.Errorf("App fiber server forced to shutdown: %v", err)
		return err
	}

	return nil
}
