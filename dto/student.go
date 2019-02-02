package dto

import (
	"github.com/satori/go.uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Age      int32     `json:"age"`
	Email    string    `json:"email"`
	Passwort string    `json:"password"`
	Address  Address   `json:"address"`
}
