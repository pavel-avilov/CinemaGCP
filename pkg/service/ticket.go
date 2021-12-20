package service

import (
	"CinemaGCP/pkg/repository"
	"github.com/google/uuid"
)

type TicketService struct {
	rep repository.Ticket
}

func NewTicketService(rep repository.Ticket) *TicketService {
	return &TicketService{rep: rep}
}

func (fs *TicketService) Buy(userId, sessionId uuid.UUID, place int) error {
	return fs.rep.Buy(userId, sessionId, place)
}
