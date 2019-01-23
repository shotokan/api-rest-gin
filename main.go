package main

import (
	"api-example/api/controllers"
	"api-example/api/routes"
	"api-example/infraestructures/gorm"
	"api-example/repositories"
	"api-example/services"
	"fmt"
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var g inject.Graph
	dbHandler := gorm.DBHandler{}
	dbHandler.Connect()
	fmt.Println(dbHandler)
	var a controllers.StudentController
	err := g.Provide(
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
	routes.RegisterRoutes(r, &a)
	r.Run()
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
