package bootstrap

import (
	"fmt"

	"github.com/mclargo/go-framework/internal/handlers"

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
	// Start the server on port 3000
	fmt.Println("Starting server in port 3000")
	if err := app.Listen(":3000"); err != nil {
		return err
	}
	// TODO: graceful shutdown

	return nil
}
