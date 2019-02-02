package models

import "time"

type BaseModel struct {
	CreatedAt time.Time `json:"-" sql:"default:now()"`
	UpdatedAt time.Time `json:"-"`
	IsActive  bool      `json:"-"`
}
