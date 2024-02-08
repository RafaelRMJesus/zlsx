package controller

import (
	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	var students []models.Student
	db.DB.Find(&students)
	c.JSON(200, students)
}
