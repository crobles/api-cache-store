package jsonValidator

import (
	jsonModel "api-cache-store/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateJson valida la estructura Cliente
func ValidateJson(cliente jsonModel.Cliente) map[string]string {
	err := validate.Struct(cliente)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.StructField()] = fmt.Sprintf("failed on the '%s' tag", err.Tag())
		}
		return errors
	}
	return nil
}



