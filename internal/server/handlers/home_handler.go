package handlers

import "net/http"

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/products", http.StatusTemporaryRedirect)
}
