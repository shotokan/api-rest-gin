package v1

import (
	"api-example/api/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.RouterGroup, student *controllers.StudentController) {
	// student := di.ServiceContainer().InjectStudentController(dbHandler)
	router.GET("/", student.GetStudent)
	router.POST("/", student.CreateStudent)
	router.GET("/:id", student.GetStudentByID)
}
