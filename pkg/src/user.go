package src

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Budget   int       `json:"budget"`
}
