package main

import (
	"lambda/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

// var w *httptest.ResponseRecorder

func init() {
	gin.SetMode(gin.ReleaseMode)
	tests.Router = SetupRuter()
}
func TestMiddleware(t *testing.T) {
	t.Run("Request ID", tests.TestRequestID)
}
func TestEndpoints(t *testing.T) {
	t.Run("Healthcheck", tests.TestHealthCheck)
}

// func Test404Route(t *testing.T) {
// 	req, _ := http.NewRequest("POST", "/", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, w.Code, http.StatusNotFound)
// }
