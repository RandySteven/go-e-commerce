package routers

import (
	"github.com/RandySteven/go-e-commerce.git/db/postgres"
	handlers "github.com/RandySteven/go-e-commerce.git/handlers/rest"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/usecases"
)

type Handlers struct {
	interfaces.UserHandler
	interfaces.ShopHandler
	interfaces.ProductHandler
}

func NewHandlers(repo *postgres.Repositories) *Handlers {
	userUsecase := usecases.NewUserUsecase(repo.UserRepository)
	shopUsecase := usecases.NewShopUsecase(repo.ShopRepository)
	productUsecase := usecases.NewProductUsecase(repo.ProductRepository, repo.ShopRepository)

	return &Handlers{
		UserHandler:    handlers.NewUserHandler(userUsecase),
		ShopHandler:    handlers.NewShopHandler(shopUsecase),
		ProductHandler: handlers.NewProductHandler(productUsecase),
	}
}
