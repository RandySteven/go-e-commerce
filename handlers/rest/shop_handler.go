package handlers_rest

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

type ShopHandler struct {
	usecase interfaces.ShopUsecase
}

// ShopDetail implements interfaces.ShopHandler.
func (*ShopHandler) ShopDetail(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx.Value("role_id"))
}

// LoginShop implements interfaces.ShopHandler.
func (h *ShopHandler) LoginShop(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(req.Context(), "request_id", requestId)
		shopLogin *requests.ShopLoginRequest
	)

	err := utils.BindJSON(req, &shopLogin)
	if err != nil {
		return
	}

	loginRes, err := h.usecase.LoginShop(ctx, shopLogin)
	if err != nil {
		return
	}

	utils.ResponseHandler(res, http.StatusOK, loginRes)
}

// RegisterShop implements interfaces.ShopHandler.
func (h *ShopHandler) RegisterShop(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var (
		requestId    = uuid.NewString()
		ctx          = context.WithValue(req.Context(), "request_id", requestId)
		shopRegister *requests.ShopRegisterRequest
	)

	err := utils.BindJSON(req, &shopRegister)
	if err != nil {
		return
	}

	shopRes, err := h.usecase.RegisterShop(ctx, shopRegister)
	if err != nil {
		log.Println(err)
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(shopRes)
}

func NewShopHandler(usecase interfaces.ShopUsecase) *ShopHandler {
	return &ShopHandler{
		usecase: usecase,
	}
}

var _ interfaces.ShopHandler = &ShopHandler{}
