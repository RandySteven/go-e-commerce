package interfaces

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, req *requests.UserRegisterRequest) (res *responses.UserResponse, err error)
		LoginUser(ctx context.Context, req *requests.UserLoginRequest) (res *responses.UserLogin, err error)
		UserDetail(ctx context.Context, id uint) (res *responses.UserDetail, err error)
	}

	ShopUsecase interface {
		RegisterShop(ctx context.Context, req *requests.ShopRegisterRequest) (res *responses.ShopResponse, err error)
		LoginShop(ctx context.Context, req *requests.ShopLoginRequest) (res *responses.ShopLoginResponse, err error)
		ShopDetail(ctx context.Context, id uint) (res *responses.ShopDetail, err error)
	}

	ProductUsecase interface {
		AddProduct(ctx context.Context, req *requests.ProductRequest) (res *models.Product, err error)
		GetAllProducts(ctx context.Context, param *query.Param) (res []models.Product, err error)
	}

	CategoryUsecase interface {
		GetAllCategories(ctx context.Context, param *query.Param) (res []models.Category, err error)
		GetCategoryById(ctx context.Context, id uint) (res *models.Category, err error)
		AddCategory(ctx context.Context, req *requests.CategoryRequest) (res *models.Category, err error)
	}
)
