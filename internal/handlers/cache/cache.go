package cacheHandler

import (
	connectDatabase "api-cache-store/database/redis"
	clientModel "api-cache-store/internal/models/jsonClient"
	jsonValidator "api-cache-store/internal/validators"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

type Response struct {
	Status   string `json:"status"`
	Mensaje string `json:"mensaje"`
	Data int `json:"data"`
}


func ValidateJson(c *fiber.Ctx) error {
	var cliente clientModel.Cliente

	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	// Validar el JSON
	if errors := jsonValidator.EvalJson(cliente); errors != nil {
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
	//-----------------------------------------------------------------------------------------------
	//funciones para el test de stress

func CreateCache(c *fiber.Ctx) error {

	var cliente2 clientModel.Cliente2
	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente2); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	jsonString, err := json.Marshal(cliente2)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to convert to JSON",
			})
	}

	DB := connectDatabase.GetRedisClient()
	err = DB.Set(ctx, cliente2.Rut, jsonString, 60 * time.Second).Err()
	if err != nil {
		fmt.Println(" ---> CREATE ERROR ❌:", err)
		//panic(err)
		return c.JSON(fiber.Map{
			"status":  "CRATE NOK",
			"mensaje": "Cache added error",
		})
	}

	fmt.Println(" ---> CREATE OK ✅:")
	return c.JSON(fiber.Map{
		"status":  "CRATE OK",
		"mensaje": "Cache added successfully",
	})
}

func ReadCache(c *fiber.Ctx) error {
	var cliente2 clientModel.Cliente2
	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente2); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	DB := connectDatabase.GetRedisClient()
	value, err := DB.Get(ctx, cliente2.Rut).Result()
	if err != nil {
		fmt.Println(" ---> READ ERROR ❌:", err)
		//panic(err)
		return c.JSON(fiber.Map{
			"status":  "READ NOK",
			"mensaje": "Cache read error",
		})
	}
	
	fmt.Println(" ---> READ OK ✅:")
	return c.JSON(fiber.Map{
		"status":  "READ OK",
		"mensaje": "Cache read successfully",
		"data": value,
	})
}

func UpdateCache(c *fiber.Ctx) error {
	var cliente2 clientModel.Cliente2
	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente2); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	now := time.Now()
	nowString := now.Format(time.RFC3339) 
	cliente2.Fecha_actualizacion = &nowString

	jsonString, err := json.Marshal(cliente2)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"mensaje": "Failed to convert to JSON",
			})
	}

	DB := connectDatabase.GetRedisClient()
	err = DB.Set(ctx, cliente2.Rut, jsonString, 60 * time.Second).Err()
	if err != nil {
		fmt.Println(" ---> UPDATE ERROR ❌:", err)
		//panic(err)
		return c.JSON(fiber.Map{
			"status":  "UPDATE NOK",
			"mensaje": "Cache updated error",
		})
	}

	fmt.Println(" ---> UPDATE OK ✅:")
	return c.JSON(fiber.Map{
		"status":  "UPDATE OK",
		"mensaje": "Cache updated successfully",
	})
}

func DeleteCache(c *fiber.Ctx) error {
	var cliente2 clientModel.Cliente2
	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente2); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	DB := connectDatabase.GetRedisClient()
	value, err := DB.Del(ctx, cliente2.Rut).Result()
	if err != nil {
		fmt.Println(" ---> DELETE ERROR ❌:", err)
		//panic(err)
		return c.JSON(fiber.Map{
			"status":  "DELETE NOK",
			"mensaje": "Cache deleted error",
		})
	}

	response := Response{}

	if value == 1 {
		fmt.Println(" ---> DELETE OK ✅:")
		response.Status = "DELETE OK"
		response.Mensaje = "Cache deleted successfully"
		response.Data = int(value)
	} else if value == 0 {
		fmt.Println(" ---> DELETE NO RESULT ✅:")
		response.Status = "DELETE NO RESULT"
		response.Mensaje = "Cache deleting not found"
		response.Data = int(value)
	}

	return c.JSON(response)
}

//-----------------------------------------------------------------------------------------------
//funciones para el test de stress con lógica

func Cache(c *fiber.Ctx) error {
	var cliente3 clientModel.Cliente3
	// Parsear el JSON recibido
	if err := c.BodyParser(&cliente3); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON WETA",
			"details": err,
		})
	}
	// buscamos el json en cache
	var rut string = cliente3.Rut
	resultGetCache := GetCache(rut)
	// verifica si el json esta en cache
	if (resultGetCache["status"] == "GET OK") {
		// verifica si el json de entrada es igual al almacenado
		if (reflect.DeepEqual(resultGetCache["data"], cliente3)) {
			return c.JSON(fiber.Map{
				"status":  "GET OK",
				"mensaje": "Cache get successfully",
				"data": resultGetCache["data"],
			})
		} else {
			// consume servicio promociones
			responseService := promotionService(cliente3)
			// actualiza cache
			responseUpdate := SetCache(responseService["data"].(clientModel.Cliente3))
			if (responseUpdate["status"] == "SET OK") {
				return c.JSON(fiber.Map{
					"status":  "SET OK",
					"mensaje": "Cache encontrada y no igual, pero se actualizo correctamente",
					"data": responseService["data"],
				})
			} else if (responseUpdate["status"] == "SET ERROR") {
				return c.JSON(fiber.Map{
					"status":  "SET NOK",
					"mensaje": "Cache add error",
				})
			}
		}
	} else if (resultGetCache["status"] == "GET NOK") {
		//consume servicio promociones
		responseService := promotionService(cliente3)
		// actualiza cache
		responseUpdate := SetCache(responseService["data"].(clientModel.Cliente3))
		if (responseUpdate["status"] == "SET OK") {
			return c.JSON(fiber.Map{
				"status":  "SET OK",
				"mensaje": "Cache no encontrada, pero se actualizo y creo correctamente",
				"data": responseService["data"],
			})
		} else if (responseUpdate["status"] == "SET ERROR") {
			return c.JSON(fiber.Map{
				"status":  "SET NOK",
				"mensaje": "Cache add error",
			})
		}
	} else {
		return c.JSON(fiber.Map{
			"status":  resultGetCache,
			"mensaje": "ERROR NO HIZO NADA",
		})
	}
	return nil // Add this return statement
}

func SetCache(dataClient clientModel.Cliente3) fiber.Map {
	jsonString, err := json.Marshal(dataClient)
	if err != nil {
			return fiber.Map{
				"status":  "SET ERROR STRINGIFY",
				"mensaje": "Failed to convert to JSON",
			}
	}

	DB := connectDatabase.GetRedisClient()
	err = DB.Set(ctx, dataClient.Rut, jsonString, 0).Err()
	if err != nil {
		fmt.Println(" ---> Error seting cache ❌:", err)
		//panic(err)
		return fiber.Map{
			"status":  "SET ERROR",
			"mensaje": "Cache setting error",
		}
	}

	fmt.Println(" ---> SETTING SUCCESSFUL ✅:")
	return fiber.Map{
		"status":  "SET OK",
		"mensaje": "Cache added successfully",
	}
}

func GetCache(rut string) fiber.Map {
	DB := connectDatabase.GetRedisClient()
	value, err := DB.Get(ctx, rut).Result()
	if err != nil {
		fmt.Println(" ---> GET NOK ❌:", err)
		//panic(err)
		return fiber.Map{
			"status":  "GET NOK",
			"mensaje": "Cache read not found",
		}
	}
	var dataJson clientModel.Cliente3
	if err := json.Unmarshal([]byte(value), &dataJson); err != nil {
		fmt.Println("Error unmarshalling JSON ", err)
		return fiber.Map{
			"status":  "GET NOK",
			"mensaje": "Error unmarshalling JSON ",
		}
	}
	fmt.Println(" ---> GET OK ✅:")
	return fiber.Map{
		"status":  "GET OK",
		"mensaje": "Cache read successfully",
		"data": dataJson,
	}
}

func promotionService(dataClient clientModel.Cliente3) fiber.Map {
	defer func() {
			if r := recover(); r != nil {
					fmt.Println(" ---> Error getting promotions: ❌", r)
			}
	}()
	
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := rng.Intn(2501) + 2000 // Tiempo aleatorio entre 2000 y 4500 ms
	time.Sleep(time.Duration(delay) * time.Millisecond)

	promotions := struct {
			Promociones []clientModel.Promocion3 `json:"promociones"`
	}{
			Promociones: []clientModel.Promocion3{
					{Id_promocion: "0h9u8t9g", Off: 0.5, Sku: "6789oikjhbgt678i"},
					{Id_promocion: "8ti56jk7uh", Off: 0.25, Sku: "5i98rf7gybhgnjk76"},
			},
	}

	dataClient.Carrito.Promociones = promotions.Promociones

	fmt.Println(" ---> GET PROMOTIONS SUCCESSFUL: ✅ ")
	return fiber.Map{
		"status":  "GET PROMOTIONS SUCCESSFUL",
		"data": dataClient,
	}
}




