package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBHandler struct {
	Conn *gorm.DB
}

func (dbh *DBHandler) Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "postgresql://ksquare:ksquare@localhost:5432/poc_chef?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// enable logger
	db.LogMode(true)
	db.DB().SetMaxOpenConns(20)
	// install extension for uuid
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	dbh.Conn = db
	return db, nil
}

func (dbh *DBHandler) DB() *gorm.DB {
	return dbh.Conn
}
