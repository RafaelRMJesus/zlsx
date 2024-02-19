package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/RafaelRMJesus/zlsx-api/controllers"
	"github.com/RafaelRMJesus/zlsx-api/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, http.StatusOK, res.Code, "expected %d, but got %d", http.StatusOK, res.Code)
	mockRes := `"Eai rafa, tudo bem?"`
	resBody, _ := io.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody))
}

func TestGetAllStudents(t *testing.T) {
	db.DbConnect()
	r := SetupTestRoutes()
	r.GET("/students", controller.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
