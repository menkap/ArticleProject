package validator

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

//ValidateInputs - Validates Inputs
func ValidateInputs(data interface{}) (bool, map[string]string) {

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		errors := make(map[string]string)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errors["err"] = "Something went wrong"
			return false, errors
		}

		reflected := reflect.TypeOf(data).Elem()

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.FieldByName(err.StructField())
			var name string
			if name = (&field).Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "The " + name + " is required"
				break
			default:
				errors[name] = "The " + name + " is invalid"
				break
			}
		}
		return false, errors
	}

	return true, nil
}
