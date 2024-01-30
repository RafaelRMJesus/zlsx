package controller

import "github.com/gin-gonic/gin"

func ShowStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "08970172012",
		"nome": "teste",
	})
}
