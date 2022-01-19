package test_utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetMockedContext(request *http.Request) (*httptest.ResponseRecorder, *gin.Context) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return response, c
}
