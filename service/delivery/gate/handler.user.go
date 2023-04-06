package gate

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/smartm2m/models"
)

func (c *Gate) UserCreate(ctx context.Context) {
	functionName := "[Gate][UserCreate]"

	request := models.CreateUserRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Request", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	user, errx := c.usecase.UserCreate(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserCreate", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		user, "Berhasil tambah pengguna")
}

func (c *Gate) UserRegularCreate(ctx context.Context) {
	functionName := "[Gate][UserRegularCreate]"

	request := models.CreateUserRequest{}

	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Request", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	userCreateResponse, errx := c.usecase.UserRegularCreate(ctx, request)

	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserRegularCreate Error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		userCreateResponse, "Berhasil tambah pengguna")
}

func (c *Gate) UserLogin(ctx context.Context) {
	functionName := "[Gate][UserLogin]"

	request := models.AuthRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Request", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	authResp, errx := c.usecase.UserLogin(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserLogin Error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		authResp, "Berhasil Login")
}

func (c *Gate) UserIndex(ctx context.Context) {
	functionName := "[Gate][UserIndex]"

	paginationData := c.utils.GetRequestPaginationData(ctx)
	search := c.utils.GetRequestQuery(ctx, "search")

	users, errx := c.usecase.UserIndexWithPagination(ctx, search, paginationData)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get User Pagination Error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		users, "Berhasil")
}

func (c *Gate) UserGetById(ctx context.Context) {
	functionName := "[Gate][UserGetById]"

	requestID := c.utils.GetRequestVar(ctx, "id")

	userId, err := strconv.Atoi(requestID)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Convert to int", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	userAndProfile, errx := c.usecase.UserGetByID(ctx, int64(userId))
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserGetByID Error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		userAndProfile, "Berhasil")
}

func (c *Gate) UserUpdatePassword(ctx context.Context) {
	functionName := "[Gate][UserUpdatePassword]"

	requestID := c.utils.GetRequestVar(ctx, "id")

	userId, err := strconv.Atoi(requestID)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Convert to int", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	request := models.CreatePasswordRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	user, errx := c.usecase.UserUpdatePassword(ctx, int64(userId), request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserUpdatePassword Error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		user, "Berhasil Sunting Password")
}

func (c *Gate) UserSelfUpdatePassword(ctx context.Context) {
	functionName := "[Gate][UserSelfUpdatePassword]"

	request := models.CreatePasswordRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	user, errx := c.usecase.UserSelfUpdatePassword(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserSelfUpdatePassword Wrong", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		user, "Berhasil Sunting Password")
}

func (c *Gate) UserUpdateImg(ctx context.Context) {
	functionName := "[Gate][UserUpdateImg]"

	requestID := c.utils.GetRequestVar(ctx, "id")

	request := models.CreateImgRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	userId, err := strconv.Atoi(requestID)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Convert to int", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	user, errx := c.usecase.UserUpdateImageURL(ctx, int64(userId), &request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While User Upload image error", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		user, "Berhasil Sunting Gambar")
}

func (c *Gate) UserUpdate(ctx context.Context) {
	functionName := "[Gate][UserUpdate]"

	requestID := c.utils.GetRequestVar(ctx, "id")

	request := models.UpdateUserRequest{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	userId, err := strconv.Atoi(requestID)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Convert to int", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	user, errx := c.usecase.UserUpdate(ctx, int64(userId), request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserUpdate", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		user, "Berhasil Sunting pengguna")
}

func (c *Gate) UserGetSession(ctx context.Context) {
	functionName := "[Gate][UserGetSession]"

	resp, errx := c.usecase.UserGetSession(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While get user Session", functionName))
		return c.utils.ResponseJSON(ctx, http.StatusUnauthorized,
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		resp, "Berhasil")
}

func (c *Gate) UserRegister(ctx context.Context) {
	functionName := "[Gate][UserRegister]"

	request := models.UserRegister{}
	if err := c.utils.BindJSONToStruct(ctx, &request); err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	err, valid, valErrors := c.utils.ValidateInputs(request)
	if err != nil {
		errx := serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	if !valid {
		errx := serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	errx := c.usecase.UserRegister(ctx, &request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Register user", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		nil, "Berhasil registrasi, silakan cek email anda untuk proses selanjutnya")
}

func (c *Gate) UserVerifyToken(ctx context.Context) {
	functionName := "[Gate][UserVerifyToken]"

	token := c.utils.GetRequestQuery(ctx, "token")

	errx := c.usecase.UserVerifyToken(ctx, token)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserVerifyToken", functionName))
		return c.utils.ResponseJSON(ctx, errx.GetStatusCode(),
			errx, errx.GetMessage())
	}

	return c.utils.ResponseJSON(ctx, http.StatusOK,
		nil, "Berhasil meregistrasi pengguna, kamu bisa login sekarang")
}
