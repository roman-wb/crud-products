package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func DataToJson(data interface{}) string {
	str, _ := json.Marshal(data)
	return string(str)
}

func BodyToString(buf *bytes.Buffer) string {
	body, _ := io.ReadAll(buf)
	return strings.TrimSpace(string(body))
}
