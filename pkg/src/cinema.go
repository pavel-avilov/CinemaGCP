package src

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Film struct {
	Id          uuid.UUID      `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description sql.NullString `json:"description" db:"description"`
	Duration    time.Time      `json:"duration" db:"duration"`
	ReleaseDate sql.NullTime   `json:"release_date" db:"release_date"`
}

type Session struct {
	Id       uuid.UUID `json:"id" db:"id"`
	IdFilm   uuid.UUID `json:"film" db:"film_id"`
	ShowDate time.Time `json:"show_date" db:"show_date"`
}

type Ticket struct {
	Id         uuid.UUID `json:"id"`
	SeatNumber int       `json:"number"`
	Price      float64   `json:"price"`
}
