package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-example/dto"
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
	var student dto.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := sc.Service.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func NewStudentController() *StudentController {
	return &StudentController{}
}
