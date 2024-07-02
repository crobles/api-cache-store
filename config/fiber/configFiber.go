package configFiber

import (
	fiberModel "api-cache-store/internal/models/fiber"
)

// estos valores tambien podrían ser configurados por medio del archiv .env
func SetConfigFiber() fiberModel.Config {
	config := fiberModel.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:	"api-cache-store v0.0.1",
	}
	return config
}
