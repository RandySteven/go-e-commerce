package usecases

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/utils"
)

type shopUsecase struct {
	shopRepo interfaces.ShopRepository
}

// LoginShop implements interfaces.ShopUsecase.
func (usecase *shopUsecase) LoginShop(ctx context.Context, req *requests.ShopLoginRequest) (res *responses.ShopLoginResponse, err error) {
	shopExists, err := usecase.shopRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// if !utils.IsPasswordValid(shopExists.Password, req.Password){
	// }
}

// RegisterShop implements interfaces.ShopUsecase.
func (usecase *shopUsecase) RegisterShop(ctx context.Context, req *requests.ShopRegisterRequest) (res *responses.ShopResponse, err error) {
	shopExists, err := usecase.shopRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if shopExists != nil {
		return nil, err
	}

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	shop := &models.Shop{
		Name:        req.Name,
		Email:       req.Email,
		Password:    pass,
		PhoneNumber: req.PhoneNumber,
	}

	shop, err = usecase.shopRepo.Save(ctx, shop)
	if err != nil {
		return nil, err
	}

	res = &responses.ShopResponse{
		Name:        shop.Name,
		Email:       shop.Email,
		PhoneNumber: shop.PhoneNumber,
	}

	return res, nil
}

// ShopDetail implements interfaces.ShopUsecase.
func (*shopUsecase) ShopDetail(ctx context.Context, id uint) (res *responses.ShopDetail, err error) {
	panic("unimplemented")
}

var _ interfaces.ShopUsecase = &shopUsecase{}

func NewShopUsecase(shopRepo interfaces.ShopRepository) *shopUsecase {
	return &shopUsecase{shopRepo: shopRepo}
}
