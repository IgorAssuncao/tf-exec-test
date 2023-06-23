package api

// import (
// 	"github.com/gofiber/fiber/v2"
// )
//
// type HealthCheck struct {
// 	Message string `json:"message"` // It has to be uppercase so it can be exported
// }
//
// func CreateServer() *fiber.App {
// 	app := fiber.New()
//
// 	app.Get("/healthcheck", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(HealthCheck{
// 			Message: "Ok",
// 		})
// 	})
//
// 	apiRouteGroup := app.Group("/api")
//
// 	apiRouteGroup.Get("/healthcheck", func(c *fiber.Ctx) error {
//    // fiber.Map allows to create dynamic json objects
//    // without having to create golang structs.
// 		return c.Status(200).JSON(fiber.Map{
// 			"message": "Ok",
// 		})
// 	})
//
// 	return app
// }
