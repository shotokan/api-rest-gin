package interfaces

import (
	"api-example/dto"
	"api-example/models"
)

type IStudentService interface {
	GetStudents()
	GetStudentById(id int32)
	CreateStudent(student *dto.Student) (*dto.Student, error)
}

type IStudentRepository interface {
	GetStudents()
	GetStudentByIdi(id int32)
	CreateStudentRepository(student *models.Student) (*models.Student, error)
}
