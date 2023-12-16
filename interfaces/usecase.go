package interfaces

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/payload/requests"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, req *requests.UserRegisterRequest)
	}
)
