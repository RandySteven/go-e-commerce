package usecases

import "github.com/RandySteven/go-e-commerce.git/interfaces"

type userUsecase struct {
	repo interfaces.UserRepository
}

var _ interfaces.UserUsecase = &userUsecase{}
