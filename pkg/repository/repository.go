package repository

import (
	"CinemaGCP/pkg/src"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user src.User) (uuid.UUID, error)
	GetUser(username, password string) (src.User, error)
	GetUserById(id uuid.UUID) (src.User, error)
}

type Film interface {
	GetAll() ([]src.Film, error)
}

type Session interface {
	GetAll() (*sql.Rows, error)
	GetSessionById(uuid.UUID) (*sql.Rows, error)
}

type Ticket interface {
	Buy(userId, sessionId uuid.UUID, place int) error
}

type Repository struct {
	Authorization
	Session
	Film
	Ticket
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Film:          NewFilmPostgres(db),
		Session:       NewSessionPostgres(db),
		Ticket:        NewTicketPostgres(db),
	}
}
