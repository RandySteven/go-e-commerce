package repositories

import (
	"context"
	"database/sql"

	"github.com/RandySteven/go-e-commerce.git/entity/models"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/scripts/gosql"
)

type userRepository struct {
	db *sql.DB
}

// GetByEmail implements interfaces.UserRepository.
func (repo *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := gosql.SelectUserByEmail
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(email).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Birthday,
			&user.PhoneNumber,
		)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete implements interfaces.UserRepository.
func (repo *userRepository) Delete(ctx context.Context, req *models.User) (res *models.User, err error) {
	panic("unimplemented")
}

// FindAll implements interfaces.UserRepository.
func (repo *userRepository) FindAll(ctx context.Context) (res []models.User, err error) {
	panic("unimplemented")
}

// FindOne implements interfaces.UserRepository.
func (repo *userRepository) FindOneById(ctx context.Context, id uint) (res *models.User, err error) {
	query := gosql.SelectUserById
	err = repo.db.QueryRowContext(ctx, query, id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save implements interfaces.UserRepository.
func (repo *userRepository) Save(ctx context.Context, req *models.User) (res *models.User, err error) {
	query := gosql.InsertUserQuery

	var userId uint
	err = repo.db.
		QueryRowContext(ctx, query, req.Name, req.Email, req.Password, req.Birthday, req.PhoneNumber).
		Scan(&userId)
	if err != nil {
		return nil, err
	}
	req.ID = userId

	return req, nil
}

// Update implements interfaces.UserRepository.
func (*userRepository) Update(ctx context.Context, req *models.User) (res *models.User, err error) {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

var _ interfaces.UserRepository = &userRepository{}
