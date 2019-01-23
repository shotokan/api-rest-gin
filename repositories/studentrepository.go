package repositories

import (
	"api-example/interfaces"
	"fmt"
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
