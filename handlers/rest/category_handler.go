package handlers_rest

import (
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/interfaces"
)

type CategoryHandler struct {
	usecase interfaces.CategoryUsecase
}

// AddCategory implements interfaces.CategoryHandler.
func (h *CategoryHandler) AddCategory(res http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

// GetAllCategories implements interfaces.CategoryHandler.
func (h *CategoryHandler) GetAllCategories(res http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

// GetCategoryById implements interfaces.CategoryHandler.
func (h *CategoryHandler) GetCategoryById(res http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

func NewCategoryHandler(usecase interfaces.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase,
	}
}

var _ interfaces.CategoryHandler = &CategoryHandler{}
