package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/RafaelRMJesus/zlsx-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestSayHello(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controller.SayHello)
	req, _ := http.NewRequest("GET", "/rafa", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("Status error: reveived %d, but expected %d", res.Code, http.StatusOK)
	}
}
