package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/DeniesKresna/smartm2m/service/repository/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	db       *sdb.DBInstance
	authCase IAuthUsecase
	repo     sql.IDatabase
}

func InitUserUsecase(db *sdb.DBInstance, repo sql.IDatabase, authCase IAuthUsecase) IUserUsecase {
	return &UserUsecase{
		db:       db,
		repo:     repo,
		authCase: authCase,
	}
}

func (h *UserUsecase) UserGetByID(ctx context.Context, ID int64) (res models.User, errx serror.SError) {
	functionName := "[UserUsecase][UserGetByID]"

	res, errx = h.repo.UserGetByID(ctx, ID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Fail to get user")
		return
	}

	return
}

func (h *UserUsecase) UserGetByEmail(ctx context.Context, email string) (res models.User, errx serror.SError) {
	functionName := "[UserUsecase][UserGetByEmail]"

	res, errx = h.repo.UserGetByEmail(ctx, email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByEmail", functionName), "Fail to get user")
		return
	}

	return
}

func (h *UserUsecase) UserCreate(ctx context.Context, payload models.UserCreatePayload) (res models.User, errx serror.SError) {
	functionName := "[UserUsecase][UserCreate]"

	_, errx = h.UserGetByEmail(ctx, payload.Email)
	if errx == nil {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusInternalServerError, fmt.Sprintf("%s User with Email %s has been exist", functionName, payload.Email), "Fail to create user")
		return
	}

	utstruct.InjectStructValue(payload, &res)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(res.Password), bcrypt.DefaultCost)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While Encrypt Password", functionName), "Fail to encrypt password")
		return
	}

	res.Password = string(hashedPassword)

	// TODO: for testing purpose, creator is system, for real case, wrap func in middleware
	// and use AuthGetFromContext instead
	creator := "system"
	res.CreatedBy = creator
	res.UpdatedBy = creator

	res, errx = h.repo.UserCreate(ctx, res)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserCreate", functionName), "Fail to create user")
		return
	}

	return
}
