package usecase

import (
	"context"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gobridge/serror"
)

type IStockUsecase interface {
	StockCreate(ctx context.Context, req models.StockCreateRequest) (s models.Stock, errx serror.SError)
	StockGetByID(ctx context.Context, id int64) (s models.Stock, errx serror.SError)
	StockBulkCreate(ctx context.Context, reqs []models.StockCreateRequest) (errx serror.SError)
}
