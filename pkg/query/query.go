package query

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	query_statement "github.com/RandySteven/go-e-commerce.git/enums"
	"github.com/RandySteven/go-e-commerce.git/utils"
)

type Repository[T any] interface {
	Save(ctx context.Context, req *T) (res *T, err error)
	FindAll(ctx context.Context) (res []T, err error)
	FindOne(ctx context.Context) (res *T, err error)
	Delete(ctx context.Context, req *T) (res *T, err error)
	Update(ctx context.Context, req *T) (res *T, err error)
}

type PostgresRepository[T any] struct {
	db *sql.DB
}

type MongoRepository[T any] struct{}

func (p *PostgresRepository[T]) Save(ctx context.Context, req *T) (res *T, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	model := utils.GetModelName(req)
	query := query_statement.INSERT_QUERY + model + "("
	values := "VALUES ("
	fields := utils.GetFields(req)
	i := 1
	var args []interface{}
	for k, v := range fields {
		query += k + ", "
		values += "$" + fmt.Sprintf("%d", i)
		i++
		args = append(args, v)
	}
	query = query[:len(query)-2] + ") "
	values = values[:len(values)-2] + ")"

	query += values

	result, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()

	selectQ := query_statement.SELECT_QUERY +
		query_statement.ALL_STATEMENT +
		query_statement.FROM_STATEMENT +
		model
	selectQ += "WHERE id = $1"
	rows, err := p.db.Query(selectQ, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(res)
		if err != nil {
			return nil, err
		}
	}

	return
}

func NewPostgresRepository[T any](db *sql.DB) *PostgresRepository[T] {
	return &PostgresRepository[T]{db: db}
}

func (m *MongoRepository[T]) Save(ctx context.Context, req *T) (res *T, err error) {
	return
}
