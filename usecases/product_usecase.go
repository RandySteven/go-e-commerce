package usecases

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type productUsecase struct {
	repo interfaces.ProductRepository
	shop interfaces.ShopRepository
}

// GetAllProducts implements interfaces.ProductUsecase.
func (usecase *productUsecase) GetAllProducts(ctx context.Context, param *query.Param) (res []models.Product, err error) {
	return usecase.repo.FindByPagination(ctx, param.Limit, param.Page)
}

// AddProduct implements interfaces.ProductUsecase.
func (usecase *productUsecase) AddProduct(ctx context.Context, req *requests.ProductRequest) (res *models.Product, err error) {
	shopEmail := ctx.Value("email").(string)
	shop, err := usecase.shop.GetByEmail(ctx, shopEmail)
	if err != nil {
		return nil, err
	}
	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		ShopID:      shop.ID,
	}
	return usecase.repo.Save(ctx, product)
}

func NewProductUsecase(repo interfaces.ProductRepository, shop interfaces.ShopRepository) *productUsecase {
	return &productUsecase{repo: repo, shop: shop}
}

var _ interfaces.ProductUsecase = &productUsecase{}
