package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func HTTPRequest(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	return w
}
