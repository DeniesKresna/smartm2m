package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
)

func FailHTTP(w http.ResponseWriter, errx serror.SError) {
	type errorData struct {
		ErrorMessage    string      `json:"error_message"`
		ErrorLine       string      `json:"error_line"`
		ErrorComment    string      `json:"error_comment"`
		ErrorValidation interface{} `json:"error_validation"`
	}

	errData := errorData{
		ErrorMessage: errx.GetErrorMessage(),
	}
	if utstring.GetEnv(models.AppENV, "dev") == "dev" {
		errData.ErrorLine = errx.GetErrorLine()
		errData.ErrorComment = errx.GetComment()
		errData.ErrorValidation = errx.GetValidation()
	}

	resp := models.ApiResponse{
		Success: false,
		Status:  errx.GetStatusCode(),
		Message: errx.GetMessage(),
		Data:    errData,
	}

	sendHTTPResponse(w, resp, errx.GetStatusCode())
	return
}

func SuccessHTTP(w http.ResponseWriter, data interface{}, message string, startTime time.Time) {
	statusCode := 200
	resp := models.ApiResponse{
		Success:     true,
		Status:      statusCode,
		Message:     message,
		ProcessTime: time.Since(startTime).Milliseconds(),
	}

	if data != nil {
		resp.Data = data
	}

	sendHTTPResponse(w, resp, statusCode)
	return
}

func sendHTTPResponse(w http.ResponseWriter, model interface{}, statusCode int) {
	json, err := json.Marshal(model)
	if err != nil {
		utlog.Error("something went wrong")
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
