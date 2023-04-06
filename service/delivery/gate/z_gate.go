package gate

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utlog"

	"github.com/go-playground/validator/v10"

	"github.com/DeniesKresna/smartm2m/models"
	"github.com/DeniesKresna/smartm2m/service/usecase"
)

type Gate struct {
	ListRoutes  []models.HTTPRoute
	Validator   *validator.Validate
	UserUsecase usecase.IUserUsecase
}

func InitGate(
	validator *validator.Validate,
	userUsecase usecase.IUserUsecase,
) *Gate {
	gate := &Gate{
		Validator:   validator,
		UserUsecase: userUsecase,
	}
	gate.InitRoutes()
	return gate
}

func (c *Gate) Get(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodGet,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Post(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))

	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPost,
		Path:    fPath,
		Handler: handl,
	}

	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Put(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPut,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Delete(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodDelete,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	var response models.ApiResponse
	response.Status = statusCode
	response.Title = message
	response.Success = false
	if statusCode < 400 {
		response.Success = true
	} else {
		errorData, ok := data.(*serror.Serror)
		if ok {
			if errorData.GetErrorMessage() != "" {
				utlog.Error(errorData.GetComment())
				utlog.Error(errorData.GetErrorLine())
			}
		}
	}
	response.Detail = data

	jsonDt, err := json.Marshal(response)
	if err != nil {
		statusCode = 500
		utlog.Error("something went wrong")
	}
	// Set CORS headers for the main request.
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDt)
}

func ValidateInputs(validate *validator.Validate, dataset interface{}) (error, bool, map[string]string) {
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
