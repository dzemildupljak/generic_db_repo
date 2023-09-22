package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterValidation("uuid", validateUUID)
}

func validateUUID(fl validator.FieldLevel) bool {
	uuidStr := fl.Field().String()
	_, err := uuid.Parse(uuidStr)
	return err == nil
}
