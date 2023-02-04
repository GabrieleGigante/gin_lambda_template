package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := HTTPRequest(req)
	var res map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, res["hello"], "lambda")
}
