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
}

type Film interface {
	GetAll() ([]src.Film, error)
}

type Sessions interface {
}

type Service struct {
	Authorization
	Sessions
	Film
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Film:          NewFilmService(rep.Film),
	}
}
