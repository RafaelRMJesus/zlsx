package controller

import (
	"fmt"
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func GetDeleted(c *gin.Context) {
	var students []models.Student
	fmt.Println(db.DB.Unscoped().Not("deleted_at", nil).Find(&students))
	c.JSON(http.StatusOK, students)
}
