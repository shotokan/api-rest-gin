package di

import (
	"api-example/api/controllers"
	"api-example/interfaces"
	"api-example/repositories"
	"api-example/services"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func NewStudentController(db *gorm.DB) (controllers.StudentController, error) {

	wire.Build(controllers.NewStudentController, services.NewStudentService, wire.Bind(new(interfaces.StudentService), repositories.StudentRepository{}))
	return controllers.StudentController{}, nil
}
