package utils

import (
	"encoding/json"
	"net/http"
)

const HeaderLocation = "Location"
const HeaderContentType = "Content-Type"
const ContentTypeJSON = "application/json"
const MessageInternalError = "Internal error"
const MessageNotFound = "Not found"

type ResponseMessage struct {
	Message string `json:"message"`
}

func ResponseOK(res http.ResponseWriter, data interface{}) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	res.WriteHeader(http.StatusOK)
	//nolint:errcheck
	json.NewEncoder(res).Encode(data)
}

func ResponseCreate(res http.ResponseWriter, data interface{}) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	res.WriteHeader(http.StatusCreated)
	//nolint:errcheck
	json.NewEncoder(res).Encode(data)
}

func ResponseNoContent(res http.ResponseWriter) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	//nolint:errcheck
	res.WriteHeader(http.StatusNoContent)
}

func ResponseInvalid(res http.ResponseWriter, data interface{}) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	res.WriteHeader(http.StatusUnprocessableEntity)
	//nolint:errcheck
	json.NewEncoder(res).Encode(data)
}

func ResponseInternalError(res http.ResponseWriter) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	res.WriteHeader(http.StatusInternalServerError)
	//nolint:errcheck
	json.NewEncoder(res).Encode(ResponseMessage{
		Message: MessageInternalError,
	})
}

func ResponseNotFound(res http.ResponseWriter) {
	res.Header().Set(HeaderContentType, ContentTypeJSON)
	res.WriteHeader(http.StatusNotFound)
	//nolint:errcheck
	json.NewEncoder(res).Encode(ResponseMessage{
		Message: MessageNotFound,
	})
}
