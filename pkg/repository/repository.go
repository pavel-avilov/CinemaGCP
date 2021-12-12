package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Sessions interface {
}

type Repository struct {
	Authorization
	Sessions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
