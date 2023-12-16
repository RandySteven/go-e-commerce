package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/query"
)

type userRepository struct {
	db *sql.DB
}

// Delete implements interfaces.UserRepository.
func (*userRepository) Delete(ctx context.Context, req *models.User) (res *models.User, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.UserRepository.
func (*userRepository) FindAll(ctx context.Context) (res []models.User, err error) {
	panic("unimplemented")
}

// FindOne implements interfaces.UserRepository.
func (*userRepository) FindOne(ctx context.Context) (res *models.User, err error) {
	panic("unimplemented")
}

// Save implements interfaces.UserRepository.
func (repo *userRepository) Save(ctx context.Context, req *models.User) (res *models.User, err error) {
	return query.NewPostgresRepository[models.User](repo.db).Save(ctx, req)
}

// Update implements interfaces.UserRepository.
func (*userRepository) Update(ctx context.Context, req *models.User) (res *models.User, err error) {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

var _ interfaces.UserRepository = &userRepository{}
