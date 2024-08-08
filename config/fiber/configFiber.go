package configFiber

import (
	fiberModel "api-cache-store/internal/models/fiber"
)

// TODO: SetConfigFiber like handler error to read ENV variables
// estos valores tambien podrían ser configurados por medio del archiv .env
func SetConfigFiber() fiberModel.Config {
	config := fiberModel.Config{
		Prefork: false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:	"api-cache-store v0.0.1",
	}
	return config
}
