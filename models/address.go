package models

import "github.com/satori/go.uuid"

type Address struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Street string    `sql:"type:varchar(200)" gorm:"not null"`
	City   string    `sql:"type:varchar(100)" gorm:"not null"`
	State  string    `sql:"type:varchar(100)" gorm:"not null"`
	BaseModel
}

func (Address) TableName() string {
	return "address"
}
