package router

import (
	configENV "api-cache-store/config/env"
	cacheRoutes "api-cache-store/internal/routes/cache"
	healthRoutes "api-cache-store/internal/routes/health"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	basePath := configENV.ConfigEnv("BASE_PATH")
	api := app.Group(basePath, logger.New())
	healthRoutes.SetupHealthRoutes(api)
	cacheRoutes.SetupCacheRoutes(api)
}