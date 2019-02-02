package interfaces

import (
	"api-example/models"
)

type IStudentService interface {
	GetStudents()
	GetStudentById(id int32)
	CreateStudent(student *models.Student)
}

type IStudentRepository interface {
	GetStudents()
	GetStudentByIdi(id int32)
	CreateStudentRepository(student *models.Student)
}
