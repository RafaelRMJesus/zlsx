package routes

import (
	controller "github.com/RafaelRMJesus/zlsx-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	r.GET("/:name", controller.SayHello)
	r.GET("/students", controller.ShowStudents)
	r.GET("/students/id/:id", controller.GetStudentById)
	r.GET("/students/name/:name", controller.SearchStudentByName)
	r.GET("/students/deleted", controller.GetDeleted)
	r.POST("/students", controller.CreateStudent)
	r.DELETE("/students/:id", controller.DeleteStudent)
	r.PATCH("/students/:id", controller.UpdateStudent)
	r.Run()
}
