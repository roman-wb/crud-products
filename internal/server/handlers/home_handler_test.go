package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roman-wb/crud-products/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_HomeHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	HomeHandler(res, req)

	assert.Equal(t, "/products", res.Header().Values(utils.HeaderLocation)[0])
	assert.Equal(t, http.StatusTemporaryRedirect, res.Result().StatusCode)
}
