package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := HTTPRequest(req)
	assert.Equal(t, res.Code, http.StatusOK)
	assert.NotEmpty(t, res.Header().Get("X-Request-ID"))
}
