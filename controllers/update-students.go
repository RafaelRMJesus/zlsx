package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	db.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Validation Error": err.Error(),
		})
		return
	}
	db.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
