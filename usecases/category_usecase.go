package usecases

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type categoryUsecase struct {
	repo interfaces.CategoryRepository
}

// AddCategory implements interfaces.CategoryUsecase.
func (*categoryUsecase) AddCategory(ctx context.Context, req *requests.CategoryRequest) (res *models.Category, err error) {
	panic("unimplemented")
}

// GetAllCategories implements interfaces.CategoryUsecase.
func (*categoryUsecase) GetAllCategories(ctx context.Context, param *query.Param) (res []models.Category, err error) {
	panic("unimplemented")
}

// GetCategoryById implements interfaces.CategoryUsecase.
func (*categoryUsecase) GetCategoryById(ctx context.Context, id uint) (res *models.Category, err error) {
	panic("unimplemented")
}

func NewCategoryUsecase(repo interfaces.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{repo: repo}
}

var _ interfaces.CategoryUsecase = &categoryUsecase{}
