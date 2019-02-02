package services

import (
	"api-example/dto"
	"api-example/interfaces"
	"api-example/models"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func (service *StudentService) CreateStudent(student *dto.Student) (*dto.Student, error) {
	log.Info("Service.CreateStudent...")
	src := models.Student{}
	src.Name = student.Name

	src.Age = student.Age
	src.IsActive = true
	src.Email = student.Email
	src.Password = student.Passwort
	add := models.Address{}
	add.City = student.Address.City
	add.State = student.Address.State
	add.Street = student.Address.Street
	add.IsActive = true
	src.Address = add
	resp, err := service.Repository.CreateStudentRepository(&src)
	if err != nil {
		log.Error("Service.CreateStudent: ", err)
		return nil, errors.Wrap(err, "error trying to create student")
	}
	student.ID = resp.ID
	student.Address.ID = resp.Address.ID
	return student, nil
}

func NewStudentService(rep interfaces.IStudentRepository) StudentService {
	return StudentService{rep}

}
