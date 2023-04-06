package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationFunction func(validator.FieldLevel) bool

func (cfg *Config) InitNewValidator() {
	v := validator.New()

	var validationFuncMap = make(map[string]ValidationFunction)

	validationFuncMap["forbiddens"] = ValidateSafeString

	ValidationRegister(v, validationFuncMap)
}

func ValidationRegister(v *validator.Validate, mvf map[string]ValidationFunction) {
	for k, vf := range mvf {
		v.RegisterValidation(k, validator.Func(vf))
	}
}

func ValidateSafeString(fl validator.FieldLevel) bool {
	enumString := fl.Param()
	value := fl.Field().String()
	enumSlice := strings.Split(enumString, "_")
	for _, v := range enumSlice {
		if strings.Contains(value, v) {
			return false
		}
	}
	return true
}
