package bootstrap

import (
	"github.com/mclargo/go-framework/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func startApp(bootstrap Bootstrap) *fiber.App {
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

	return app
}
