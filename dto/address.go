package dto

import "github.com/satori/go.uuid"

type Address struct {
	ID     uuid.UUID `json:"id"`
	Street string    `json:"street"`
	City   string    `json:"city"`
	State  string    `json:"state"`
}
