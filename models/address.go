package models

import "time"

type Address struct {
	Street    string
	City      string
	State     string
	UpdatedAt time.Time
	BaseModel
}
