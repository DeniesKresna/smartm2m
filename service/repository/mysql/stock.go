package mysql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/danatest/service/repository/mysql/queries"
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
)

type MysqlStockRepository struct {
	db *sdb.DBInstance
}

func InitMysqlStockRepository(sdb *sdb.DBInstance) IMysqlStockRepository {
	return &MysqlStockRepository{
		db: sdb,
	}
}

func (r *MysqlStockRepository) StockCreate(ctx context.Context, req models.StockCreateRequest) (id int64, errx serror.SError) {
	functionName := "[MysqlRepository][StockCreate]"

	res, err := r.db.Exec(queries.QueryCreateStock,
		req.Name, req.Price, req.Availability, req.IsActive,
	)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While QueryCreateStock (data: %+v)", functionName, req))
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Last Inserted ID", functionName))
		return
	}
	return
}

func (r *MysqlStockRepository) StockGetByID(ctx context.Context, id int64) (s models.Stock, errx serror.SError) {
	functionName := "[MysqlRepository][StockGetByID]"

	err := r.db.Take(&s, queries.QueryGetStockByID, id)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While QueryGetStockByID (id: %+v)", functionName, id))
		return
	}

	return
}
