package usecase

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/danatest/service/repository/mysql/mocks"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mysqlStockRepository = &mocks.IMysqlStockRepository{Mock: mock.Mock{}}
	stockUsecase         = &StockUsecase{stockMysql: mysqlStockRepository}

	timeNow = time.Now()

	// error list
	errNotFound   = serror.NewWithErrorComment(errors.New("not found"), http.StatusNotFound, "stock not found")
	errCreateFail = serror.NewWithErrorComment(errors.New("create fail"), http.StatusNotFound, "stock cannot be created")

	// response list
	stock1 = models.Stock{
		ID:           1,
		Name:         "Tas",
		Price:        200000,
		Availability: 1,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}
)

func TestStockGetByID(t *testing.T) {
	var (
		listIds = []int64{0, 1}
		ctx     = context.Background()
	)

	// mock definition
	mysqlStockRepository.Mock.On("StockGetByID", ctx, listIds[0]).Return(models.Stock{}, errNotFound)
	mysqlStockRepository.Mock.On("StockGetByID", ctx, listIds[1]).Return(stock1, nil)
	// end of mock definition

	stockOne := models.Stock{
		ID:           1,
		Name:         "Tas",
		Price:        200000,
		Availability: 1,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}

	testCases := []models.TestCases{
		{
			Title:    "Get Stock Success",
			Args:     []interface{}{ctx, int64(1)},
			Expected: []interface{}{stockOne, nil},
		}, {
			Title:    "Get Stock Fail not found",
			Args:     []interface{}{ctx, int64(0)},
			Expected: []interface{}{models.Stock{}, errNotFound},
		},
	}

	for _, v := range testCases {
		stock, errx := stockUsecase.StockGetByID(v.Args[0].(context.Context), v.Args[1].(int64))
		if v.Expected[0] == nil {
			assert.Nil(t, stock)
		} else {
			assert.Equal(t, v.Expected[0].(models.Stock), stock)
		}

		if v.Expected[1] == nil {
			assert.Nil(t, errx)
		} else {
			assert.Equal(t, v.Expected[1].(serror.SError), errx)
		}
	}
}

func TestStockCreate(t *testing.T) {
	var (
		listIds = []int64{0, 1}

		availabiltyTotal    = 1
		listStockCreateArgs = []models.StockCreateRequest{
			{
				Name:         "Tas",
				Price:        200000,
				Availability: &availabiltyTotal,
			}, {
				Name:         "Buku",
				Price:        0,
				Availability: &availabiltyTotal,
			},
		}
		ctx = context.Background()
	)

	// mock definition
	mysqlStockRepository.Mock.On("StockGetByID", ctx, listIds[0]).Return(models.Stock{}, errNotFound)
	mysqlStockRepository.Mock.On("StockGetByID", ctx, listIds[1]).Return(stock1, nil)

	mysqlStockRepository.Mock.On("StockCreate", ctx, listStockCreateArgs[0]).Return(int64(1), nil)
	mysqlStockRepository.Mock.On("StockCreate", ctx, listStockCreateArgs[1]).Return(int64(0), errCreateFail)
	// end of mock definition

	stockOne := models.Stock{
		ID:           1,
		Name:         "Tas",
		Price:        200000,
		Availability: 1,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}

	testCases := []models.TestCases{
		{
			Title:    "Create Stock Success",
			Args:     []interface{}{ctx, listStockCreateArgs[0]},
			Expected: []interface{}{stockOne, nil},
			Message:  "Create Stock Success",
		}, {
			Title:    "Create Stock Fail",
			Args:     []interface{}{ctx, listStockCreateArgs[1]},
			Expected: []interface{}{models.Stock{}, errCreateFail},
			Message:  "Price should be greater than 0",
		},
	}

	for _, v := range testCases {
		stockID, errx := stockUsecase.StockCreate(v.Args[0].(context.Context), v.Args[1].(models.StockCreateRequest))
		if v.Expected[0] == nil {
			assert.Nil(t, stockID)
		} else {
			assert.Equal(t, v.Expected[0].(models.Stock), stockID, v.Message)
		}

		if v.Expected[1] == nil {
			assert.Nil(t, errx)
		} else {
			assert.Equal(t, v.Expected[1].(serror.SError), errx)
		}
	}
}
