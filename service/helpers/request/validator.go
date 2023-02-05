package request

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateInputs(dataset interface{}) (error, bool, map[string]string) {
	validate := validator.New()
	errors := make(map[string]string)
	err := validate.Struct(dataset)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			return err, false, errors
		}
		datasetPtr := reflect.ValueOf(dataset)
		datasetVal := reflect.Indirect(datasetPtr)
		datasetType := datasetVal.Type()

		for _, Valerr := range err.(validator.ValidationErrors) {
			field, _ := datasetType.FieldByName(Valerr.StructField())
			name := Valerr.StructField()
			errors[name] = field.Tag.Get("valerr")
		}

		return nil, false, errors
	}
	return nil, true, errors
}
