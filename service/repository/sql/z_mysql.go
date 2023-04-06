package sql

import (
	"context"

	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/myqgen2/qgen"

	"github.com/DeniesKresna/smartm2m/models"
)

type DatabaseRepository struct {
	db *sdb.DBInstance
	q  *qgen.Obj
}

func InitDatabase(db *sdb.DBInstance, q *qgen.Obj) IDatabase {
	return &DatabaseRepository{
		db: db,
		q:  q,
	}
}

type IDatabase interface {
	//user
	UserCreate(ctx context.Context, userPayload models.User) (res models.User, errx serror.SError)
	UserGetByID(ctx context.Context, ID int64) (res models.User, errx serror.SError)
	UserGetByEmail(ctx context.Context, email string) (res models.User, errx serror.SError)

	//item
	ItemsGetByUserID(ctx context.Context, userID int64) (res []models.Item, errx serror.SError)
	ItemGetByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError)
	ItemCreate(ctx context.Context, item models.ItemCreatePayload) (res models.Item, errx serror.SError)
	ItemUpdateByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError)
	ItemDeleteByID(ctx context.Context, ID int64) (errx serror.SError)
	ItemPurchasedByID(ctx context.Context, ID int64, userID int64) (errx serror.SError)
}
