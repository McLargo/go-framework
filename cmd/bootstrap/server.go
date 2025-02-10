package bootstrap

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/mclargo/go-framework/internal/handlers"
	"golang.org/x/sync/errgroup"

	"github.com/gofiber/fiber/v2"
)

func RunServer() error {
	// TODO: init config

	// TODO: init logger

	// initialize a new Fiber app
	app := fiber.New()

	// set endpoint -> healthz
	app.Get("/healthz", handlers.Healthz)

	// TODO: set dynamic port in conf

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer done()

	ggroup, gctx := errgroup.WithContext(ctx)

	// Start the server on port 3000
	ggroup.Go(func() error {
		if err := app.Listen(":3000"); err != nil && err != context.Canceled {
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
