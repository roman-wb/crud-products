package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Response_Const(t *testing.T) {
	require.Equal(t, "Location", HeaderLocation)
	require.Equal(t, "Content-Type", HeaderContentType)
	require.Equal(t, "application/json", ContentTypeJSON)
	require.Equal(t, "Internal error", MessageInternalError)
	require.Equal(t, "Not found", MessageNotFound)
}

func Test_ResponseOK(t *testing.T) {
	// given
	data := struct {
		Status string
	}{
		Status: "ok",
	}
	res := httptest.NewRecorder()

	// when
	ResponseOK(res, data)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusOK, res.Result().StatusCode)
	require.Equal(t, DataToJson(data), BodyToString(res.Body))
}

func Test_ResponseCreate(t *testing.T) {
	// given
	data := struct {
		Status string
	}{
		Status: "ok",
	}
	res := httptest.NewRecorder()

	// when
	ResponseCreate(res, data)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusCreated, res.Result().StatusCode)
	require.Equal(t, DataToJson(data), BodyToString(res.Body))
}

func Test_ResponseNoContent(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseNoContent(res)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusNoContent, res.Result().StatusCode)
	require.Equal(t, "", BodyToString(res.Body))
}

func Test_ResponseInvalid(t *testing.T) {
	// given
	data := struct {
		Status string
	}{
		Status: "ok",
	}
	res := httptest.NewRecorder()

	// when
	ResponseInvalid(res, data)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode)
	require.Equal(t, DataToJson(data), BodyToString(res.Body))
}

func Test_ResponseInternalError(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseInternalError(res)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	require.Equal(t, DataToJson(ResponseMessage{
		Message: MessageInternalError,
	}), BodyToString(res.Body))
}

func Test_ResponseNotFound(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseNotFound(res)

	// then
	require.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	require.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	require.Equal(t, DataToJson(ResponseMessage{
		Message: MessageNotFound,
	}), BodyToString(res.Body))
}
