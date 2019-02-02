package main

import (
	"api-example/infrastructures/gorm"
	"api-example/models"
	"log"
)

func main() {
	dbHandler := &gorm.DBHandler{}
	err := dbHandler.Connect()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	Migration(dbHandler)
}

func Migration(conn *gorm.DBHandler) {
	student := new(models.Student)
	address := new(models.Address)
	conn.DB().DropTableIfExists(address, student)
	conn.DB().CreateTable(address)
	conn.DB().CreateTable(student)
	conn.DB().Model(student).AddIndex("idx_student_name", "name")
	conn.DB().Model(student).AddForeignKey("address_id", "\"address\"(id)", "RESTRICT", "RESTRICT")
}
