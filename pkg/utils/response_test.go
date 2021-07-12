package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Response_Const(t *testing.T) {
	// then
	assert.Equal(t, "Location", HeaderLocation)
	assert.Equal(t, "Content-Type", HeaderContentType)
	assert.Equal(t, "application/json", ContentTypeJSON)
	assert.Equal(t, "Internal error", MessageInternalError)
	assert.Equal(t, "Not found", MessageNotFound)
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
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, DataToJson(data), BodyToString(res.Body))
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
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusCreated, res.Result().StatusCode)
	assert.Equal(t, DataToJson(data), BodyToString(res.Body))
}

func Test_ResponseNoContent(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseNoContent(res)

	// then
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusNoContent, res.Result().StatusCode)
	assert.Equal(t, "", BodyToString(res.Body))
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
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode)
	assert.Equal(t, DataToJson(data), BodyToString(res.Body))
}

func Test_ResponseInternalError(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseInternalError(res)

	// then
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, DataToJson(ResponseMessage{
		Message: MessageInternalError,
	}), BodyToString(res.Body))
}

func Test_ResponseNotFound(t *testing.T) {
	// given
	res := httptest.NewRecorder()

	// when
	ResponseNotFound(res)

	// then
	assert.Equal(t, ContentTypeJSON, res.Header().Values(HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, DataToJson(ResponseMessage{
		Message: MessageNotFound,
	}), BodyToString(res.Body))
}
