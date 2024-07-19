package router

import (
	configENV "api-cache-store/config/env"
	cacheRoutes "api-cache-store/internal/routes/cache"
	healthRoutes "api-cache-store/internal/routes/health"
	cacheRoutes_stressTest "api-cache-store/internal/routes/stressTest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	basePath := configENV.ConfigEnv("BASE_PATH")
	api := app.Group(basePath, logger.New())
	healthRoutes.SetupHealthRoutes(api)
	cacheRoutes.SetupCacheRoutes(api)
	//parte del test de estress
	cacheRoutes_stressTest.SetupCacheRoutes_stressTest(api)
}