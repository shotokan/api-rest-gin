package routes

import (
	"api-example/api/controllers"
	"api-example/api/routes/v1"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, student *controllers.StudentController) {
	apiV1 := r.Group("/v1")

	v1.StudentRoutes(apiV1.Group("/students"), student)
}
