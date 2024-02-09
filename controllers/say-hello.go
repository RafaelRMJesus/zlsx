package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, "Eai "+name+", tudo bem?")
}
