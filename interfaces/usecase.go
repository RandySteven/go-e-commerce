package interfaces

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
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
)
