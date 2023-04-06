package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/DeniesKresna/smartm2m/service/repository/sql"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	db   *sdb.DBInstance
	repo sql.IDatabase
}

func InitAuthUsecase(db *sdb.DBInstance, repo sql.IDatabase) IAuthUsecase {
	return &AuthUsecase{
		db:   db,
		repo: repo,
	}
}

func (h *AuthUsecase) AuthGetFromContext(ctx context.Context) (res models.User, errx serror.SError) {
	functionName := "[AuthUsecase][AuthGetFromContext]"

	session := ctx.Value("session")
	sessionType, ok := session.(models.Session)
	if !ok {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusUnauthorized, fmt.Sprintf("%s Session Not Found", functionName), "Session not Found")
		return
	}
	userID := sessionType.UserID
	if userID <= 0 {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusUnauthorized, fmt.Sprintf("%s Session Not Found", functionName), "Session not Found")
		return
	}

	res, errx = h.repo.UserGetByID(ctx, userID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While User SessionGetByID", functionName), "Fail to get user session")
		return
	}

	return
}

func (h *AuthUsecase) AuthLogin(ctx context.Context, email string, password string) (res models.TokenResponse, errx serror.SError) {
	functionName := "[AuthUsecase][AuthGetFromContext]"

	user, errx := h.repo.UserGetByEmail(ctx, email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByEmail", functionName), "User not found")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusMethodNotAllowed, "[Usecase][UserLogin] While CompareHashAndPassword", "Password is wrong")
		return
	}

	// token generation
	{
		expires := time.Now().Add(time.Hour * 24)

		// Create the JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
			Issuer:    "my-app",
			Subject:   fmt.Sprintf("%d", user.ID),
		})

		// Sign the token with a secret key
		tokenString, err := token.SignedString([]byte(utstring.GetEnv(models.AppApiSecret)))
		if err != nil {
			errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While SignedString token", functionName), "Cannot generate token")
			return
		}

		// Set the token response
		res.Token = tokenString
		res.Expires = expires
	}

	return
}
