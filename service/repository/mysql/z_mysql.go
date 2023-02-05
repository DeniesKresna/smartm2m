package mysql

import (
	"context"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/gobridge/serror"
)

type IMysqlStockRepository interface {
	StockCreate(ctx context.Context, req models.StockCreateRequest) (id int64, errx serror.SError)
	StockGetByID(ctx context.Context, id int64) (s models.Stock, errx serror.SError)
}
