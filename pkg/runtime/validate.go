package runtime

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// TODO: Add xml support
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

var validate *validator.Validate

func ValidateInput(params interface{}) validator.ValidationErrors {
	if params == nil {
		return nil
	}

	err := validate.Struct(params)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return validationErrors
	}

	return nil
}
