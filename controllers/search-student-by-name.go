package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func SearchStudentByName(c *gin.Context) {
	var student models.Student
	name := c.Params.ByName("name")
	db.DB.Where(&models.Student{Name: name}).First(&student)
	if student.Name == "" {
		c.JSON(http.StatusNotFound, "Student Not Found")
		return
	}
	c.JSON(http.StatusOK, student)
}
