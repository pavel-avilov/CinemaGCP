package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SessionPostgres struct {
	db *sqlx.DB
}

func NewSessionPostgres(db *sqlx.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (fp *SessionPostgres) GetAll() (*sql.Rows, error) {
	query := fmt.Sprintf(
		`SELECT s.id, f.name, f.duration, s.show_date FROM %v s
		INNER JOIN film f on f.id = s.film_id`, sessionTable)
	rows, err := fp.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (fp *SessionPostgres) GetSessionById(id uuid.UUID) (*sql.Rows, error) {
	query := fmt.Sprintf(
		`SELECT s.id, f.name, f.duration, s.show_date, u_s.place FROM %v s
		INNER JOIN film f on f.id = s.film_id
		LEFT JOIN user_session u_s on u_s.session_id = s.id
		WHERE s.id = $1`, sessionTable)
	rows, err := fp.db.Query(query, id.String())
	if err != nil {
		return nil, err
	}
	return rows, nil
}
