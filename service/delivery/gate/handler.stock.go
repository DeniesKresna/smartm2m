package gate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/danatest/service/helpers/request"
	"github.com/DeniesKresna/danatest/service/helpers/response"
	"github.com/DeniesKresna/gobridge/serror"
)

func (c *Gate) AddStock(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = r.Context()
		functionName = "[Handler][AddStock]"
		startTime    = time.Now()
		message      = "Berhasil"
	)

	var stockCreateRequest models.StockCreateRequest
	err := json.NewDecoder(r.Body).Decode(&stockCreateRequest)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Gagal memproses masukan")
		response.FailHTTP(w, errx)
		return
	}

	err, valid, valErrors := request.ValidateInputs(stockCreateRequest)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		response.FailHTTP(w, errx)
		return
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s Input Payload is wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		response.FailHTTP(w, errx)
		return
	}

	stock, errx := c.StockUsecase.StockCreate(ctx, stockCreateRequest)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While StockCreate", functionName))
		response.FailHTTP(w, errx)
		return
	}

	response.SuccessHTTP(w, stock, message, startTime)
	return
}

func (c *Gate) GetStockByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = r.Context()
		functionName = "[Handler][GetStockByID]"
		startTime    = time.Now()
		message      = "Berhasil"
		stock        models.Stock
	)

	id := request.GetInt64Var(r, "id")

	stock, errx := c.StockUsecase.StockGetByID(ctx, id)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While StockCreate", functionName))
		response.FailHTTP(w, errx)
		return
	}

	response.SuccessHTTP(w, stock, message, startTime)
	return
}

func (c *Gate) CreateStockBulk(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = r.Context()
		functionName = "[Handler][CreateStockBulk]"
		startTime    = time.Now()
		message      = "Berhasil"
		payloads     []models.StockCreateRequest
		errx         serror.SError
	)

	err := json.NewDecoder(r.Body).Decode(&payloads)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(
			err, http.StatusInternalServerError, fmt.Sprintf("%s While json.NewDecoder", functionName), "Gagal mengkonversi masukan",
		)
		response.FailHTTP(w, errx)
		return
	}

	for _, payload := range payloads {
		err, valid, valErrors := request.ValidateInputs(payload)
		if err != nil {
			errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
			response.FailHTTP(w, errx)
			return
		}

		if !valid {
			errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s Input Payload is wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
			response.FailHTTP(w, errx)
			return
		}
	}

	errx = c.StockUsecase.StockBulkCreate(ctx, payloads)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While StockBulkCreate", functionName))
		response.FailHTTP(w, errx)
		return
	}

	response.SuccessHTTP(w, nil, message, startTime)
	return
}
