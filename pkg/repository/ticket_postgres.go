package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TicketPostgres struct {
	db *sqlx.DB
}

func NewTicketPostgres(db *sqlx.DB) *TicketPostgres {
	return &TicketPostgres{db: db}
}

func (fp *TicketPostgres) Buy(userId, sessionId uuid.UUID, place int) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id, session_id, place)  VALUES ($1, $2, $3)`, ticketTable)
	_, err := fp.db.Exec(query, userId, sessionId, place)
	if err != nil {
		return err
	}
	return nil
}
