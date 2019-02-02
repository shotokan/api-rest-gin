package repositories

import (
	"api-example/interfaces"
	"api-example/models"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type StudentRepository struct {
	Conn interfaces.IDbHandler `inject:""`
}

func (sr *StudentRepository) GetStudents() {

}
func (sr *StudentRepository) GetStudentByIdi(id int32) {
	fmt.Println(10)
	fmt.Println(sr.Conn.DB())

}

func (sr *StudentRepository) CreateStudentRepository(student *models.Student) {
	fmt.Println("Creating student")
	log.Info("Creating student")
	err := sr.Conn.DB().Create(&student).Error

	if err != nil {
		log.Error(err)
		fmt.Println("Hubo un error: ", err)
	}
}
