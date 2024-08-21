package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthzResponse struct {
	Server string `json:"server"`
}

func Healthz(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(&HealthzResponse{"OK"})
}
