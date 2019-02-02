package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-example/interfaces"
	"api-example/models"
)

type StudentController struct {
	Service interfaces.IStudentService `inject:""`
}

func (sc *StudentController) GetStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"profile": "estudiante"})
}

func (sc *StudentController) GetStudentByID(c *gin.Context) {
	src := models.Student{}
	src.Name = "Ivan"

	sc.Service.GetStudentById(10)
	src.Age = 32
	src.Email = "asdasd@asdsad.com"
	src.Password = "12345678"

	c.JSON(http.StatusOK, src)
}

func (sc *StudentController) CreateStudent(c *gin.Context) {
	src := models.Student{}
	src.Name = "Ivan"

	src.Age = 32
	src.Email = "asdasd@asdsad.com"
	src.Password = "12345678"
	add := models.Address{}
	add.City = "Merida"
	add.State = "Yucatan"
	add.Street = "calle 63"
	add.IsActive = true
	src.Address = add
	sc.Service.CreateStudent(&src)

	c.JSON(http.StatusOK, src)
}

func NewStudentController() *StudentController {
	return &StudentController{}
}
