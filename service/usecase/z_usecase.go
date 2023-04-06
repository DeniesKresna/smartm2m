package usecase

import (
	"context"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/smartm2m/models"
)

type IUserUsecase interface {
	//user
	UserCreate(ctx context.Context, payload models.UserCreatePayload) (res models.User, errx serror.SError)
	UserGetByEmail(ctx context.Context, email string) (res models.User, errx serror.SError)
	UserGetByID(ctx context.Context, ID int64) (res models.User, errx serror.SError)
}

type IAuthUsecase interface {
	//auth
	AuthGetFromContext(ctx context.Context) (res models.User, errx serror.SError)
}

type IItemUsecase interface {
	//item
	ItemsGetByUserID(ctx context.Context, userID int64) (res []models.Item, errx serror.SError)
	ItemGetByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError)
	ItemCreate(ctx context.Context, item models.Item) (errx serror.SError)
	ItemUpdateByID(ctx context.Context, ID int64) (res models.Item, errx serror.SError)
	ItemDeleteByID(ctx context.Context, ID int64) (errx serror.SError)
	ItemPurchasedByID(ctx context.Context, ID int64, userID int64) (errx serror.SError)
}
