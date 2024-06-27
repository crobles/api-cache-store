package cacheHandler

import (
	jsonModel "api-cache-store/internal/models"
	jsonValidator "api-cache-store/internal/validators"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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