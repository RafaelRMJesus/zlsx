package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says": "Hello " + name,
	})
}
