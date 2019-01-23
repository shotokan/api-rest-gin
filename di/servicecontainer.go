package di

import (
	"api-example/api/controllers"
	"api-example/repositories"
	"api-example/services"
	"sync"

	db "api-example/infraestructures/gorm"
)

type IServiceContainer interface {
	InjectStudentController(dbHandler *db.DBHandler) controllers.StudentController
}

type kernel struct{}

func (k *kernel) InjectStudentController(dbHandler *db.DBHandler) controllers.StudentController {

	studentRepository := &repositories.StudentRepository{dbHandler}
	studentService := &services.StudentService{studentRepository}
	studentController := controllers.StudentController{studentService}

	return studentController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
