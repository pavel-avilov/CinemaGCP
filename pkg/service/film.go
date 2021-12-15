package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
)

type FilmService struct {
	rep repository.Film
}

func NewFilmService(rep repository.Film) *FilmService {
	return &FilmService{rep: rep}
}

func (fs *FilmService) GetAll() ([]src.Film, error) {
	return fs.rep.GetAll()
}
