package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"test.com/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// https://circleci.com/blog/gin-gonic-testing/
func TestGroupSensorGeneration(t *testing.T) {
	r := SetUpRouter()
	var GroupController = new(controllers.GroupController)
	r.GET("api/auth/", GroupController.GroupSensorGeneration)
	req, _ := http.NewRequest("GET", "/api/auth/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

//
func TestDataGeneration(t *testing.T) {
	r := SetUpRouter()
	var GroupController = new(controllers.GroupController)
	r.GET("api/auth/", GroupController.DataGeneration)
	req, _ := http.NewRequest("GET", "/api/auth/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
