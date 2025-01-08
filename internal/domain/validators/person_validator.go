package validators

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"personCrud/internal/domain/models"
)

type PersonValidator struct {
	validate *validator.Validate
}

// NewPersonValidator crea una nueva instancia del validador
func NewPersonValidator() *PersonValidator {
	return &PersonValidator{validate: validator.New()}
}

// Validate valida el modelo Person
func (pv *PersonValidator) Validate(person *models.Person) error {
	fmt.Print("Validating person")
	err := pv.validate.Struct(person)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "ID":
					return fmt.Errorf("ID must be a numeric value and is required")
				case "Name":
					return fmt.Errorf("Name must be a string between 3 and 50 characters")
				case "Address":
					return fmt.Errorf("Address must be a string between 5 and 100 characters")
				case "Phone":
					if fieldError.Tag() == "numeric" {
						return fmt.Errorf("Phone must be a numeric value")
					}
					return fmt.Errorf("Phone number must be exactly 10 digits")
				case "Email":
					if fieldError.Tag() == "email" {
						return fmt.Errorf("Email must be a valid email address")
					}
					return fmt.Errorf("Email is required")
				case "CreatedAt", "UpdatedAt":
					return fmt.Errorf("%s must be a valid datetime", fieldError.Field())
				default:
					return fmt.Errorf("%s is invalid", fieldError.Field())
				}
			}
		}
		return fmt.Errorf("validation failed: %v", err)
	}
	return nil
}
