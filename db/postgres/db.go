package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/pkg/configs"
	"github.com/RandySteven/go-e-commerce.git/repositories"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Repositories struct {
	interfaces.AddressRepository
	interfaces.UserRepository
	db *sql.DB
}

func NewRepositories(config *configs.Config) (*Repositories, error) {
	postgresConn := config.Postgres
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresConn.Host,
		postgresConn.Port,
		postgresConn.DbUser,
		postgresConn.DbPass,
		postgresConn.DbName,
	)
	log.Println(conn)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		AddressRepository: repositories.NewAddressRepository(db),
		UserRepository:    repositories.NewUserRepository(db),
		db:                db,
	}, nil
}
