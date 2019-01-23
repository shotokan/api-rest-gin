package interfaces

import "github.com/jinzhu/gorm"

type IDbHandler interface {
	DB() *gorm.DB
	Connect() (*gorm.DB, error)
}
