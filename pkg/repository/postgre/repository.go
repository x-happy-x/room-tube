package postgre

import (
	"github.com/jmoiron/sqlx"
	"tube/pkg/config"
	"tube/pkg/repository"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(dbConfig config.Database) (repository.Repository, error) {
	db, err := sqlx.Connect("postgres", dbConfig.GetDataSource())
	return &Repository{
		DB: db,
	}, err
}
