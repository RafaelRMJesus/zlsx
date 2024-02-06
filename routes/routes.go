package routes

import (
	controller "github.com/RafaelRMJesus/zlsx-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	r.GET("/students", controller.ShowStudents)
	r.GET("/hello/:name", controller.Hello)
	r.Run()
}
