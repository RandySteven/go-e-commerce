package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/scripts/gosql"
)

type categoryRepository struct {
	db *sql.DB
}

// Delete implements interfaces.CategoryRepository.
func (repo *categoryRepository) Delete(ctx context.Context, req *models.Category) (res *models.Category, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.CategoryRepository.
func (repo *categoryRepository) FindAll(ctx context.Context) (res []models.Category, err error) {
	query := gosql.SelectCategories
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	category := &models.Category{}
	for rows.Next() {
		rows.Scan(
			&category.ID,
			&category.Category,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		res = append(res, *category)
	}
	return res, nil
}

// FindOneById implements interfaces.CategoryRepository.
func (repo *categoryRepository) FindOneById(ctx context.Context, id uint) (res *models.Category, err error) {
	panic("unimplemented")
}

// Save implements interfaces.CategoryRepository.
func (repo *categoryRepository) Save(ctx context.Context, req *models.Category) (res *models.Category, err error) {
	query := gosql.InsertCategoryQuery
	var categoryId uint = 0
	err = repo.db.QueryRowContext(ctx, query, &req.Category).Scan(&categoryId)
	if err != nil {
		return nil, err
	}
	req.ID = categoryId
	return req, nil
}

// Update implements interfaces.CategoryRepository.
func (repo *categoryRepository) Update(ctx context.Context, req *models.Category) (res *models.Category, err error) {
	panic("unimplemented")
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

var _ interfaces.CategoryRepository = &categoryRepository{}
