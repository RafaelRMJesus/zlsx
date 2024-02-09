package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func GetDeleted(c *gin.Context) {
	var students []models.Student
	db.DB.Where("deleted_at IS NULL").Find(&students)
	c.JSON(http.StatusOK, students)
}
