package controllers

import (
	"net/http"
	"strconv"

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
	id := c.Param("id")
	src := models.Student{}
	src.Name = "Ivan"
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		panic(err)
	}
	sc.Service.GetStudentById(10)
	src.A.B = int32(i)
	src.Age = 32
	src.Email = "asdasd@asdsad.com"
	src.Password = "12345678"

	c.JSON(http.StatusOK, src)
}

func NewStudentController() *StudentController {
	return &StudentController{}
}
