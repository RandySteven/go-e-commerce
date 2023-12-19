package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
)

type shopRepository struct {
	db *sql.DB
}

// Delete implements interfaces.ShopRepository.
func (repo *shopRepository) Delete(ctx context.Context, req *models.Shop) (res *models.Shop, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.ShopRepository.
func (repo *shopRepository) FindAll(ctx context.Context) (res []models.Shop, err error) {
	panic("unimplemented")
}

// FindOne implements interfaces.ShopRepository.
func (repo *shopRepository) FindOne(ctx context.Context) (res *models.Shop, err error) {
	panic("unimplemented")
}

// GetByEmail implements interfaces.ShopRepository.
func (repo *shopRepository) GetByEmail(ctx context.Context, email string) (*models.Shop, error) {
	panic("unimplemented")
}

// Save implements interfaces.ShopRepository.
func (repo *shopRepository) Save(ctx context.Context, req *models.Shop) (res *models.Shop, err error) {
	query := "INSERT INTO shops (name, email, password, phone_number) VALUES " +
		"($1, $2, $3, $4)"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var shopId uint
	err = stmt.QueryRowContext(ctx,
		req.Name,
		req.Email,
		req.Password,
		req.PhoneNumber,
	).Scan(&shopId)

	if err != nil {
		return nil, err
	}
	req.ID = shopId

	return req, nil
}

// Update implements interfaces.ShopRepository.
func (repo *shopRepository) Update(ctx context.Context, req *models.Shop) (res *models.Shop, err error) {
	panic("unimplemented")
}

func NewShopRepository(db *sql.DB) *shopRepository {
	return &shopRepository{db: db}
}

var _ interfaces.ShopRepository = &shopRepository{}
