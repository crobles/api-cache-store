package cacheHandler

import (
	connectDatabase "api-cache-store/database/redis"
	jsonModel "api-cache-store/internal/models"
	jsonValidator "api-cache-store/internal/validators"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func ValidateJson(c *fiber.Ctx) error {
	var cliente jsonModel.Cliente

		// Parsear el JSON recibido
		if err := c.BodyParser(&cliente); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		// Validar el JSON
		if errors := jsonValidator.ValidateJson(cliente); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		// Imprimir el JSON
		jsonString, _ := json.MarshalIndent(cliente, "", "  ")
		fmt.Println(string(jsonString))

		// Retornar el JSON recibido
		return c.JSON(cliente)
	}

	func CacheAdd(c *fiber.Ctx) error {
		
		// Ejemplo de uso del cliente Redis
		DB := connectDatabase.GetRedisClient()
		err := DB.Set(ctx, "key3", "value", 60 * time.Second).Err()
		if err != nil {
			fmt.Println(" error ---- :", err)
			//panic(err)
		}
		return c.SendString("CacheAdd")
	}

	func CacheGet(c *fiber.Ctx) error {
		DB := connectDatabase.GetRedisClient()
		value, err := DB.Get(ctx, "key3").Result()
		if err != nil {
			fmt.Println("error:", err)
			//panic(err)
		}
		fmt.Println("Valor almacenado en 'key':", value)
		return c.SendString("CacheGet")
	}

