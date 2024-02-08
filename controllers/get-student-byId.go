package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	db.DB.Find(&student, id)
	if student.Name == "" {
		c.JSON(http.StatusNotFound, "Student Not Found")
		return
	}
	c.JSON(http.StatusOK, student)
}
