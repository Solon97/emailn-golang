package validator

import (
	internalerrors "emailn/internal/errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	validationError := err.(validator.ValidationErrors)[0]
	field := strings.ToLower(validationError.Field())
	switch validationError.Tag() {
	case "required":
		return fmt.Errorf(internalerrors.ErrRequiredFieldPattern, field)
	case "min":
		return fmt.Errorf(internalerrors.ErrMinFieldPattern, field, validationError.Param())
	case "max":
		return fmt.Errorf(internalerrors.ErrMaxFieldPattern, field, validationError.Param())
	case "email":
		return fmt.Errorf(internalerrors.ErrEmailFieldPattern, field)
	}

	return err
}
