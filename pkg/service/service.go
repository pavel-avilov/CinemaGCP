package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(user src.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
}

type Sessions interface {
}

type Service struct {
	Authorization
	Sessions
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
