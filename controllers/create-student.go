package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
