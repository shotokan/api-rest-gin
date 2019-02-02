package main

import (
	"api-example/api/controllers"
	"api-example/api/routes"
	"api-example/infrastructures/gorm"
	"api-example/infrastructures/logs"
	"api-example/repositories"
	"api-example/services"
	"fmt"
	"log"

	"api-example/models"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var g inject.Graph
	dbHandler := gorm.DBHandler{}
	err := dbHandler.Connect()
	fmt.Println(dbHandler)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	Migrate(&dbHandler)
	logs.NewLog()
	var a controllers.StudentController
	var ro routes.MyRoutes
	err = g.Provide(
		&inject.Object{Value: &dbHandler},
		&inject.Object{Value: &repositories.StudentRepository{}},
		&inject.Object{Value: &services.StudentService{}},
		&inject.Object{Value: &a},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := g.Populate(); err != nil {
		log.Fatal(err)
	}
	// adding controllers to routes
	ro.Student = a
	ro.RegisterRoutes(r)
	r.Run(":3000")
}

func Migrate(dbh *gorm.DBHandler) {
	dbh.Conn.AutoMigrate(&models.Address{}, &models.Student{})
	dbh.Conn.Model(&models.Student{}).AddForeignKey("address_id", "address(id)", "RESTRICT", "RESTRICT")

}

// type MyInterface interface {
// 	Foo()
// }

// type Implementation struct{}

// func (imp *Implementation) Foo() {
// 	fmt.Println("Hello")
// }

// type StructDependingOnInterface struct {
// 	Bar MyInterface `inject:""`
// }

// func main() {
// 	var g inject.Graph

// 	s := &StructDependingOnInterface{}
// 	err := g.Provide(
// 		&inject.Object{Value: s},
// 		&inject.Object{Value: &Implementation{}},
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := g.Populate(); err != nil {
// 		log.Fatal(err)
// 	}

// 	s.Foo()
// }
