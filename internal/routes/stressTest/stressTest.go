package cacheRoutes_stressTest

import (
	cacheHandler "api-cache-store/internal/handlers/cache"

	"github.com/gofiber/fiber/v2"
)


func SetupCacheRoutes_stressTest(router fiber.Router) {
	cache := router.Group("/stressTestCache")
	// pruebda de cache con crud
	cache.Post("/create", cacheHandler.CreateCache)
	cache.Get("/read", cacheHandler.ReadCache)
	cache.Put("/update", cacheHandler.UpdateCache)
	cache.Delete("/delete", cacheHandler.DeleteCache)
	// prueba de cache con logica
	cache.Post("/cache", cacheHandler.Cache)
}