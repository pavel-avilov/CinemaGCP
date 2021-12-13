package repository

import (
	"CinemaGCP/pkg/src"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user src.User) (uuid.UUID, error)
	GetUser(username, password string) (src.User, error)
}

type Sessions interface {
}

type Repository struct {
	Authorization
	Sessions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
