package models

import "github.com/satori/go.uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `sql:"type:varchar(150)" gorm:"not null"`
	Age       int32     `sql:"type:smallint" gorm:"not null"`
	Email     string    `sql:"type:varchar(150)" gorm:"not null"`
	Password  string    `sql:"type:varchar(150)" gorm:"not null"`
	Address   Address   `gorm:"foreignkey:AddressID"`
	AddressID uuid.UUID `sql:"type:uuid"`
}
