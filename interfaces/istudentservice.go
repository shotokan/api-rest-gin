package interfaces

type IStudentService interface {
	GetStudents()
	GetStudentById(id int32)
}

type IStudentRepository interface {
	GetStudents()
	GetStudentByIdi(id int32)
}
