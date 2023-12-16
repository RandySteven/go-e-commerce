package interfaces

import (
	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type (
	UserRepository interface {
		query.Repository[models.User]
	}
)
