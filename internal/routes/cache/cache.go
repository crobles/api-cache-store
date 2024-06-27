package cacheRoutes

import (
	cacheHandler "api-cache-store/internal/handlers/cache"

	"github.com/gofiber/fiber/v2"
)


func SetupCacheRoutes(router fiber.Router) {
	cache := router.Group("/cache")

	cache.Post("/cacheStore", cacheHandler.ValidateJson)
}