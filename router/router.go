package router

import (
	config "api-cache-store/config"
	cacheRoutes "api-cache-store/internal/routes/cache"
	healthRoutes "api-cache-store/internal/routes/health"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var basePath string = config.Config("BASE_PATH")

func SetupRoutes(app *fiber.App) {
	api := app.Group(basePath, logger.New())
	healthRoutes.SetupHealthRoutes(api)
	cacheRoutes.SetupCacheRoutes(api)

}