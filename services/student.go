package services

import (
	"api-example/interfaces"
	"fmt"
)

type StudentService struct {
	Repository interfaces.IStudentRepository `inject:""`
}

func (service *StudentService) GetStudents() {
	fmt.Println("En get students service")
}

func (service *StudentService) GetStudentById(id int32) {
	fmt.Println("En GetStudeintById")
	service.Repository.GetStudentByIdi(id)
}

func NewStudentService(rep interfaces.IStudentRepository) StudentService {
	return StudentService{rep}

}
