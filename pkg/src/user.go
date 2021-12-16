package src

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"-" db:"id"`
	Name     string    `json:"name" binding:"required" db:"name"`
	Username string    `json:"username" binding:"required" db:"username"`
	Password string    `json:"password" binding:"required" db:"password_hash"`
	Budget   int       `json:"budget" db:"budget"`
}
