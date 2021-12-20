package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(src.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(string) (uuid.UUID, error)
	GetUser(uuid.UUID) (user src.User, err error)
}

type Film interface {
	GetAll() ([]src.Film, error)
}

type Session interface {
	GetAll() ([]sessionInfoList, error)
	GetSessionById(uuid.UUID) (SessionInfo, error)
}

type Ticket interface {
	Buy(userId, sessionId uuid.UUID, place int) error
}

type Service struct {
	Authorization
	Session
	Film
	Ticket
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Film:          NewFilmService(rep.Film),
		Session:       NewSessionService(rep.Session),
		Ticket:        NewTicketService(rep.Ticket),
	}
}
