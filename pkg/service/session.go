package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/utils"
	"database/sql"
	"github.com/google/uuid"
)

type SessionService struct {
	rep repository.Session
}

func NewSessionService(rep repository.Session) *SessionService {
	return &SessionService{rep: rep}
}

type sessionInfoList struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Duration string    `json:"duration"`
	ShowDate string    `json:"show_date"`
}

func (fs *SessionService) GetAll() ([]sessionInfoList, error) {
	var sessionData []sessionInfoList
	sessionRows, err := fs.rep.GetAll()
	if err != nil {
		return nil, err
	}
	for sessionRows.Next() {
		var s sessionInfoList
		err = sessionRows.Scan(&s.Id, &s.Name, &s.Duration, &s.ShowDate)
		if err != nil {
			return nil, err
		}
		sessionData = append(sessionData, s)
	}
	return sessionData, nil
}

type SessionInfo struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Duration       string    `json:"duration"`
	ShowDate       string    `json:"show_date"`
	AvailableSeats []int     `json:"available_seats"`
}

func (fs *SessionService) GetSessionById(id uuid.UUID) (SessionInfo, error) {
	var (
		info        SessionInfo
		takenPlaces []int
	)

	sessionRows, err := fs.rep.GetSessionById(id)
	if err != nil {
		return SessionInfo{}, err
	}
	for sessionRows.Next() {
		var (
			place sql.NullInt32
		)
		err = sessionRows.Scan(&info.Id, &info.Name, &info.Duration, &info.ShowDate, &place)
		if err != nil {
			return SessionInfo{}, err
		}
		if place.Valid {
			takenPlaces = append(takenPlaces, int(place.Int32))
		}
	}
	info.AvailableSeats = utils.Difference(utils.GenerateAvailableSeats(50), takenPlaces)
	return info, nil
}
