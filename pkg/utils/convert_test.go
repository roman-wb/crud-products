package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DataToJson(t *testing.T) {
	tesCases := []struct {
		name string
		data interface{}
		want string
	}{
		{
			name: "struct",
			data: struct {
				Status string  `json:"status"`
				Number float64 `json:"number"`
			}{
				Status: "ok",
				Number: 100.99,
			},
			want: `{"status":"ok","number":100.99}`,
		},
		{
			name: "string",
			data: "as string",
			want: `"as string"`,
		},
		{
			name: "number",
			data: 100.99,
			want: "100.99",
		},
		{
			name: "null",
			data: nil,
			want: `null`,
		},
	}

	for _, tc := range tesCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := DataToJson(tc.data)
			require.Equal(t, tc.want, got)
		})
	}
}

func Test_BodyToString(t *testing.T) {
	tesCases := []struct {
		name string
		data string
		want string
	}{
		{
			name: "string",
			data: "   data   \n    ",
			want: "data",
		},
		{
			name: "new line",
			data: "\n",
			want: "",
		},
		{
			name: "blank",
			data: "",
			want: "",
		},
	}

	for _, tc := range tesCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := BodyToString(bytes.NewBufferString(tc.data))
			require.Equal(t, tc.want, got)
		})
	}
}
