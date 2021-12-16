package service

import (
	"CinemaGCP/pkg/repository"
	"github.com/google/uuid"
)

type SessionService struct {
	rep repository.Session
}

func NewSessionService(rep repository.Session) *SessionService {
	return &SessionService{rep: rep}
}

type sessionInfo struct {
	id       uuid.UUID
	name     string
	duration string
	showDate string
}

func (fs *SessionService) GetAll() ([]map[string]string, error) {
	var sessionData []map[string]string
	sessionRows, err := fs.rep.GetAll()
	if err != nil {
		return nil, err
	}
	for sessionRows.Next() {
		var s sessionInfo
		err = sessionRows.Scan(&s.id, &s.name, &s.duration, &s.showDate)
		if err != nil {
			return nil, err
		}
		sessionData = append(sessionData, map[string]string{
			"id":           s.id.String(),
			"filmName":     s.name,
			"filmDuration": s.duration,
			"showData":     s.showDate,
		})
	}
	return sessionData, nil
}
