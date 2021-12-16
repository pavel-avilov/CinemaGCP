package repository

import (
	"database/sql"
	"fmt"
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
