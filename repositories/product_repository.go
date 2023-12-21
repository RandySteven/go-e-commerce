package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
)

type productRepository struct {
	db *sql.DB
}

// Delete implements interfaces.ProductRepository.
func (repo *productRepository) Delete(ctx context.Context, req *models.Product) (res *models.Product, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.ProductRepository.
func (repo *productRepository) FindAll(ctx context.Context) (res []models.Product, err error) {
	query := "SELECT id, name, price, stock, description, shop_id, created_at, updated_at FROM products"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
}

// FindOne implements interfaces.ProductRepository.
func (repo *productRepository) FindOneById(ctx context.Context, id uint) (res *models.Product, err error) {
	query := "SELECT * FROM products WHERE id = $1"
	err = repo.db.QueryRowContext(ctx, query, id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements interfaces.ProductRepository.
func (repo *productRepository) Save(ctx context.Context, req *models.Product) (res *models.Product, err error) {
	query := "INSERT INTO products (name, price, stock, description, shop_id, category_id) " +
		" VALUES " +
		"($1, $2, $3, $4, $5, $6) RETURNING ID"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var productId uint = 0
	err = stmt.QueryRow(
		req.Name,
		req.Price,
		req.Stock,
		req.Description,
		req.ShopID,
		req.CategoryID,
	).
		Scan(&productId)

	if err != nil {
		return nil, err
	}

	return req, nil
}

// Update implements interfaces.ProductRepository.
func (*productRepository) Update(ctx context.Context, req *models.Product) (res *models.Product, err error) {
	panic("unimplemented")
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

var _ interfaces.ProductRepository = &productRepository{}
