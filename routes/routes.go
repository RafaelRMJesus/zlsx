package routes

import (
	controller "github.com/RafaelRMJesus/zlsx-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	r.GET("/students", controller.ShowStudents)
	r.POST("/students", controller.CreateStudent)
	r.GET("/students/:id", controller.GetStudentById)
	r.DELETE("students/:id", controller.DeleteStudent)
	r.Run()
}
