package server

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/roman-wb/crud-products/internal/repos"
	"github.com/stretchr/testify/require"
)

func Test_NewRouter(t *testing.T) {
	testCases := []struct {
		name   string
		method string
		query  string
		want   bool
	}{
		{
			method: "GET",
			query:  "/",
			want:   true,
		},
		{
			method: "POST",
			query:  "/",
			want:   false,
		},
		{
			method: "GET",
			query:  "/health",
			want:   true,
		},
		{
			method: "POST",
			query:  "/health",
			want:   false,
		},
		{
			method: "GET",
			query:  "/products",
			want:   true,
		},
		{
			method: "POST",
			query:  "/products",
			want:   true,
		},
		{
			method: "PATCH",
			query:  "/products",
			want:   true,
		},
		{
			method: "PUT",
			query:  "/products",
			want:   true,
		},
		{
			method: "DELETE",
			query:  "/products",
			want:   false,
		},
		{
			method: "GET",
			query:  "/products/1",
			want:   true,
		},
		{
			method: "POST",
			query:  "/products/1",
			want:   true,
		},
		{
			method: "PATCH",
			query:  "/products/1",
			want:   true,
		},
		{
			method: "PUT",
			query:  "/products/1",
			want:   true,
		},
		{
			method: "DELETE",
			query:  "/products/1",
			want:   true,
		},
	}

	router := NewRouter(nil, repos.NewRepos(nil))

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("Request %s %s should be %v", tc.method, tc.query, tc.want)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(tc.method, tc.query, nil)
			m := &mux.RouteMatch{}
			got := router.Match(req, m)

			require.Equal(t, tc.want, got)
		})
	}
}
