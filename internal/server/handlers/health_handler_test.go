package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roman-wb/crud-products/pkg/utils"
	"github.com/stretchr/testify/require"
)

func Test_HealthHandler(t *testing.T) {
	res := httptest.NewRecorder()

	HealthHandler(res, nil)

	require.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	require.Equal(t, http.StatusOK, res.Result().StatusCode)
	require.Equal(t, utils.DataToJson(ResponseHealth{
		Status: "ok",
	}), utils.BodyToString(res.Body))
}
