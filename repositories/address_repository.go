package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type addressRepository struct {
	db *sql.DB
}

// Delete implements interfaces.AddressRepository.
func (*addressRepository) Delete(ctx context.Context, req *models.Address) (res *models.Address, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.AddressRepository.
func (*addressRepository) FindAll(ctx context.Context) (res []models.Address, err error) {
	panic("unimplemented")
}

// FindOne implements interfaces.AddressRepository.
func (*addressRepository) FindOneById(ctx context.Context, id uint) (res *models.Address, err error) {
	panic("unimplemented")
}

// Save implements interfaces.AddressRepository.
func (repo *addressRepository) Save(ctx context.Context, req *models.Address) (res *models.Address, err error) {
	return query.NewPostgresRepository[models.Address](repo.db).Save(ctx, req)
}

// Update implements interfaces.AddressRepository.
func (*addressRepository) Update(ctx context.Context, req *models.Address) (res *models.Address, err error) {
	panic("unimplemented")
}

var _ interfaces.AddressRepository = &addressRepository{}

func NewAddressRepository(db *sql.DB) *addressRepository {
	return &addressRepository{db: db}
}
