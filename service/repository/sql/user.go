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

func (r *DatabaseRepository) UserGetByID(ctx context.Context, ID int64) (u models.User, errx serror.SError) {
	functionName := "[DatabaseRepository][UserGetByID]"

	err := r.db.Take(&u, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"user*",
		},
		Conditions: map[string]interface{}{
			"id": ID,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserGetByID (id: %d)", functionName, ID))
		return
	}
	return
}

func (r DatabaseRepository) UserCreate(ctx context.Context, user models.User) (res models.User, errx serror.SError) {
	functionName := "[DatabaseRepository][UserCreate]"

	qRes, err := r.db.Exec(queries.InsertUser,
		user.CreatedBy,
		user.UpdatedBy,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Password,
	)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query InsertUser (user: %+v)", functionName, user))
		return
	}

	lastID, err := qRes.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Inserted User (user: %+v)", functionName, user))
		return
	}

	res, errx = r.UserGetByID(ctx, lastID)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserGetByID (userId: %d)", functionName, lastID))
		return
	}
	return
}

func (r DatabaseRepository) UserGetByEmail(ctx context.Context, email string) (u models.User, errx serror.SError) {
	functionName := "[DatabaseRepository][UserGetByEmail]"

	err := r.db.Take(&u, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"user*",
		},
		Conditions: map[string]interface{}{
			"email": email,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserGetByEmail (email: %s)", functionName, email))
		return
	}
	return
}
