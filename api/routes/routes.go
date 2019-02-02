package routes

import (
	"api-example/api/controllers"
	"api-example/api/routes/v1"

	"github.com/gin-gonic/gin"
)

type MyRoutes struct {
	Student controllers.StudentController
}

func (mr MyRoutes) RegisterRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")

	v1.StudentRoutes(apiV1.Group("/students"), &mr.Student)
}
