package service

import "CinemaGCP/pkg/repository"

type Authorization interface {
}

type Sessions interface {
}

type Service struct {
	Authorization
	Sessions
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
