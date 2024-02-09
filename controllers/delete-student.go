package controller

import (
	"net/http"

	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/RafaelRMJesus/zlsx-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	db.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"Student Suscessfully deleted": student})
}
