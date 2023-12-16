package usecases

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
)

type userUsecase struct {
	userRepo    interfaces.UserRepository
	addressRepo interfaces.AddressRepository
}

// LoginUser implements interfaces.UserUsecase.
func (*userUsecase) LoginUser(ctx context.Context, req *requests.UserLoginRequest) (res *responses.UserLogin, err error) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserUsecase.
func (*userUsecase) RegisterUser(ctx context.Context, req *requests.UserRegisterRequest) (res *responses.UserResponse, err error) {
	panic("unimplemented")
}

// UserDetail implements interfaces.UserUsecase.
func (*userUsecase) UserDetail(ctx context.Context, id uint) (res *responses.UserDetail, err error) {
	panic("unimplemented")
}

func NewUserUsecase(userRepo interfaces.UserRepository, addressRepo interfaces.AddressRepository) *userUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
