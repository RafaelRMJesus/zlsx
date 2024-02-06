package controller

import (
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}
