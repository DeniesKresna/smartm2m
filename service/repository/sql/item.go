package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/myqgen2/qgen"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/DeniesKresna/smartm2m/service/repository/sql/queries"
)

func (r *DatabaseRepository) ItemsGetByUserID(ctx context.Context, userID int64) (res []models.Item, errx serror.SError) {
	functionName := "[DatabaseRepository][ItemsGetByUserID]"

	qArgs := qgen.Args{
		Fields: []string{
			"item*",
		},
		Sorting: []string{"-id"},
	}

	if userID > 0 {
		qArgs.Conditions["user_id"] = userID
	}

	err := r.db.Select(&res, r.q.Build(queries.GetItems, qArgs))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query GetItems", functionName))
		return
	}
	return
}
func (r *DatabaseRepository) ItemGetByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError) {
	functionName := "[DatabaseRepository][ItemGetByID]"

	err := r.db.Take(&res, r.q.Build(queries.GetItems, qgen.Args{
		Fields: []string{
			"item*",
		},
		Conditions: map[string]interface{}{
			"id": ID,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetItem By ID. (id:%d)", functionName, ID))
		return
	}
	return
}
func (r *DatabaseRepository) ItemCreate(ctx context.Context, item models.ItemCreatePayload) (res models.Item, errx serror.SError) {
	functionName := "[DatabaseRepository][ItemCreate]"

	qRes, err := r.db.Exec(queries.InsertItem, item.UserID, item.Data)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query Insert Item", functionName))
		return
	}

	lastID, err := qRes.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Inserted Item (item: %+v)", functionName, item))
		return
	}

	res, errx = r.ItemGetByID(ctx, lastID)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While ItemGetByID (itemID: %d)", functionName, lastID))
		return
	}

	return
}
func (r *DatabaseRepository) ItemUpdateByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError) {
	return
}
func (r *DatabaseRepository) ItemDeleteByID(ctx context.Context, ID int64) (errx serror.SError) {
	return
}
func (r *DatabaseRepository) ItemPurchasedByID(ctx context.Context, ID int64, userID int64) (errx serror.SError) {
	return
}
