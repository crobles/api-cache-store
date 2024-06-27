package healthRoutes

import (
	"github.com/gofiber/fiber/v2"
)


func SetupHealthRoutes(router fiber.Router) {
	health := router.Group("/health")

	health.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("The API is UP!")
		return err
	})
}