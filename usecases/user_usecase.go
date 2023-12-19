package usecases

import (
	"context"
	"time"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/auth"
	"github.com/RandySteven/go-e-commerce.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, req *requests.UserLoginRequest) (res *responses.UserLogin, err error) {
	user, err := usecase.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	isValid := utils.IsPasswordValid(user.Password, req.Password)
	if !isValid {
		return nil, err
	}

	expTime := time.Now().Add(time.Hour * 1)
	claims := &auth.JWTClaim{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APPLICATION",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(auth.JWT_KEY)
	if err != nil {
		return nil, err
	}
	res = &responses.UserLogin{
		Token: token,
	}
	return res, nil
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, req *requests.UserRegisterRequest) (res *responses.UserResponse, err error) {
	birthdate, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, err
	}

	err = utils.DateValidation(birthdate)
	if err != nil {
		return nil, err
	}

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:        req.FirstName + req.LastName,
		Email:       req.Email,
		Password:    pass,
		PhoneNumber: req.PhoneNumber,
		Birthday:    birthdate,
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
