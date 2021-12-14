package src

import (
	"github.com/google/uuid"
	"time"
)

type Film struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    time.Time `json:"duration"`
	ReleaseDate time.Time `json:"date_of_published"`
}

type Session struct {
	Id       uuid.UUID `json:"id"`
	IdFilm   uuid.UUID `json:"film"`
	ShowDate time.Time `json:"show_date"`
}

type Ticket struct {
	Id         uuid.UUID `json:"id"`
	SeatNumber int       `json:"number"`
	Price      float64   `json:"price"`
}
