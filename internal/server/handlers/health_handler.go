package handlers

import (
	"net/http"

	"github.com/roman-wb/crud-products/pkg/utils"
)

type ResponseHealth struct {
	Status string `json:"status"`
}

func HealthHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseOK(res, ResponseHealth{
		Status: "ok",
	})
}
