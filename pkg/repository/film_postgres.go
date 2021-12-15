package repository

import (
	"CinemaGCP/pkg/src"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FilmPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *FilmPostgres {
	return &FilmPostgres{db: db}
}

func (fp *FilmPostgres) GetAll() ([]src.Film, error) {
	var films []src.Film
	query := fmt.Sprintf(`SELECT id, name, description, duration, release_date FROM %s`, filmTable)
	if err := fp.db.Select(&films, query); err != nil {
		return nil, err
	}

	return films, nil
}
