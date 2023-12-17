package usecases

import (
	"context"
	"time"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
}

// LoginUser implements interfaces.UserUsecase.
func (*userUsecase) LoginUser(ctx context.Context, req *requests.UserLoginRequest) (res *responses.UserLogin, err error) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, req *requests.UserRegisterRequest) (res *responses.UserResponse, err error) {
	//Should add email validation
	date, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:        req.FirstName + req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Birthday:    date,
	}
	user, err = usecase.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	res = &responses.UserResponse{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Birthdate:   user.Birthday.Format("2006-01-02"),
	}
	return res, nil
}

// UserDetail implements interfaces.UserUsecase.
func (*userUsecase) UserDetail(ctx context.Context, id uint) (res *responses.UserDetail, err error) {
	panic("unimplemented")
}

func NewUserUsecase(userRepo interfaces.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
