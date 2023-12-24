package interfaces

import (
	"context"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type (
	UserRepository interface {
		query.Repository[models.User]
		GetByEmail(ctx context.Context, email string) (*models.User, error)
	}

	ShopRepository interface {
		query.Repository[models.Shop]
		GetByEmail(ctx context.Context, email string) (*models.Shop, error)
	}

	AddressRepository interface {
		query.Repository[models.Address]
	}

	ShopAddressRepository interface {
		query.Repository[models.ShopAddress]
	}

	UserAddressRepository interface {
		query.Repository[models.UserAddress]
	}

	ProductRepository interface {
		query.Repository[models.Product]
		FindByPagination(ctx context.Context, limitItem uint, page uint) ([]models.Product, error)
	}

	TransactionRepository interface {
		PurchaseTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	}

	CartRepository interface {
		InsertToCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	}

	CategoryRepository interface {
		query.Repository[models.Category]
	}
)
