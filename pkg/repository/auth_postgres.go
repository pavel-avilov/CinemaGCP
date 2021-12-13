package repository

import (
	"CinemaGCP/pkg/logger"
	"CinemaGCP/pkg/src"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user src.User) (uuid.UUID, error) {
	var id uuid.UUID
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, budget) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := ap.db.QueryRow(query, user.Name, user.Username, user.Password, user.Budget)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (src.User, error) {
	var user src.User
	logger.Info(user.Username)
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := ap.db.Get(&user, query, username, password)
	logger.Info(user.Username)
	return user, err
}
