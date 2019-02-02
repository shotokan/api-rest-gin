package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	User
	Grade int `sql:"type:smallint" gorm:"not null"`
	BaseModel
}

func (Student) TableName() string {
	return "student"
}

func (s *Student) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(s.Password), 12)
	fmt.Println("generating hash...")
	if err != nil {
		return errors.Wrap(err, "could not generate the hash for password")
	}
	scope.SetColumn("Password", string(hash))
	return nil
}
