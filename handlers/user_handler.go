package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/enums/content_type"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/utils"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var (
		requestId    = uuid.NewString()
		ctx          = context.WithValue(req.Context(), "request_id", requestId)
		loginRequest *requests.UserLoginRequest
	)
	err := utils.BindJSON(req, &loginRequest)
	if err != nil {
		return
	}

	loginRes, err := handler.usecase.LoginUser(ctx, loginRequest)
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loginRes)
}

// RegisterUser implements interfaces.UserHandler.
func (handler *UserHandler) RegisterUser(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var (
		requestId    = uuid.NewString()
		ctx          = context.WithValue(req.Context(), "request_id", requestId)
		userRegister *requests.UserRegisterRequest
	)
	err := utils.BindJSON(req, &userRegister)
	if err != nil {
		log.Fatal(err)
		return
	}

	userRes, err := handler.usecase.RegisterUser(ctx, userRegister)
	if err != nil {
		log.Fatal(err)
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(userRes)
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
