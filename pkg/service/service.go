package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(user src.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uuid.UUID, error)
	GetUser(userId uuid.UUID) (user src.User, err error)
}

type Film interface {
	GetAll() ([]src.Film, error)
}

type Session interface {
	GetAll() ([]map[string]string, error)
}

type Service struct {
	Authorization
	Session
	Film
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Film:          NewFilmService(rep.Film),
		Session:       NewSessionService(rep.Session),
	}
}
