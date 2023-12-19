package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/enums/content_type"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/utils"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var loginRequest *requests.UserLoginRequest
	err := utils.BindJSON(req, &loginRequest)
	if err != nil {
		return
	}

	loginRes, err := handler.usecase.LoginUser(context.Background(), loginRequest)
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loginRes)
}

// RegisterUser implements interfaces.UserHandler.
func (handler *UserHandler) RegisterUser(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var userRegister *requests.UserRegisterRequest
	err := utils.BindJSON(req, &userRegister)
	if err != nil {
		return
	}

	userRes, err := handler.usecase.RegisterUser(context.Background(), userRegister)
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(userRes)
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
