package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roman-wb/crud-products/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_HealthHandler(t *testing.T) {
	res := httptest.NewRecorder()

	HealthHandler(res, nil)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(ResponseHealth{
		Status: "ok",
	}), utils.BodyToString(res.Body))
}
