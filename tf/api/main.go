package api

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheck struct {
	Message string `json:"message"`
}

func CreateServer() *fiber.App {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(HealthCheck{
			Message: "Ok",
		})
	})

	apiRouteGroup := app.Group("/api")

	apiRouteGroup.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Message": "Ok",
		})
	})

	return app
}
