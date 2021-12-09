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
	Id               uuid.UUID   `json:"id"`
	IdFilm           uuid.UUID   `json:"film"`
	IdHall           uuid.UUID   `json:"id_hall"`
	ShowDate         time.Time   `json:"show_date"`
	AvailableTickets []uuid.UUID `json:"available_tickets"`
}

type Hall struct {
	Id          uuid.UUID `json:"id"`
	SeatsNumber int       `json:"seats_nubmer"`
}

type Ticket struct {
	Id         uuid.UUID `json:"id"`
	SeatNumber int       `json:"number"`
	Price      float64   `json:"price"`
}
