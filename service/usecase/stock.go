package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/danatest/service/repository/mysql"
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
)

type StockUsecase struct {
	db         *sdb.DBInstance
	stockMysql mysql.IMysqlStockRepository
}

func InitStockUsecase(db *sdb.DBInstance, stockMysql mysql.IMysqlStockRepository) IStockUsecase {
	return &StockUsecase{
		db:         db,
		stockMysql: stockMysql,
	}
}

func (u *StockUsecase) StockCreate(ctx context.Context, req models.StockCreateRequest) (s models.Stock, errx serror.SError) {
	functionName := "[StockUsecase][StockCreate]"

	//handle tx
	err := u.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		u.db.SubmitTx(err)
	}()

	id, errx := u.stockMysql.StockCreate(ctx, req)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While StockCreate", functionName), "Gagal tambah data Stock")
		return
	}

	s, errx = u.stockMysql.StockGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While StockGetByID", functionName), "Gagal ambil data Stock")
		return
	}

	return
}

func (u *StockUsecase) StockGetByID(ctx context.Context, id int64) (s models.Stock, errx serror.SError) {
	functionName := "[StockUsecase][StockGetByID]"

	s, errx = u.stockMysql.StockGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While StockGetByID", functionName), "Gagal ambil data Stock")
		return
	}

	return
}
