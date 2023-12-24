package handlers

import (
	"context"
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/enums/content_type"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
	"github.com/RandySteven/go-e-commerce.git/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	usecase interfaces.ProductUsecase
}

// CreateProduct implements interfaces.ProductHandler.
func (h *ProductHandler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	utils.ContentType(res, content_type.ApplicationJson)
	var (
		requestId  = uuid.NewString()
		ctx        = context.WithValue(req.Context(), "request_id", requestId)
		productReq *requests.ProductRequest
	)

	err := utils.BindJSON(req, &productReq)
	if err != nil {
		return
	}

	product, err := h.usecase.AddProduct(ctx, productReq)
	if err != nil {
		return
	}

	utils.ResponseHandler(res, http.StatusCreated, product)
}

// GetAllProducts implements interfaces.ProductHandler.
func (h *ProductHandler) GetAllProducts(res http.ResponseWriter, req *http.Request) {
	paramQuery := mux.Vars(req)
	var (
		par       = query.DefaultParam()
		requestId = uuid.NewString()
		ctx       = context.WithValue(req.Context(), "request_id", requestId)
	)
	if paramQuery != nil {
		par.SortBy = paramQuery["sort_by"]
		par.Sort = paramQuery["sort"]
	}

	products, err := h.usecase.GetAllProducts(ctx, par)
	if err != nil {
		return
	}

	utils.ResponseHandler(res, http.StatusOK, products)
}

var _ interfaces.ProductHandler = &ProductHandler{}

func NewProductHandler(usecase interfaces.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}
