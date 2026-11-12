package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"interseguro/matrix-api/internal/adapters/statistics"
)

func TestAnalysis(t *testing.T) {
	for _, tc := range []struct {
		name, operation, body   string
		upstream, status, count int
	}{
		{"qr", "qr", "", 200, 200, 2}, {"rotate", "rotate", "", 200, 200, 1},
		{"upstream error", "qr", "", 500, 502, 2}, {"timeout", "qr", "", 0, 502, 2},
		{"broken response", "qr", "{}", 200, 502, 2}, {"invalid operation", "x", "", 500, 422, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			stub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				if r.Method != "POST" || r.URL.Path != "/internal/v1/statistics" || r.Header.Get("X-Request-ID") != "test" {
					t.Error("invalid upstream request")
				}
				var body struct{ Matrices [][][]float64 }
				if json.NewDecoder(r.Body).Decode(&body) != nil || len(body.Matrices) != tc.count {
					t.Error("invalid matrices")
				}
				if tc.operation == "rotate" && body.Matrices[0][0][0] != 3 {
					t.Error("expected rotated matrix")
				}
				if tc.upstream == 0 {
					<-r.Context().Done()
					return
				}
				w.WriteHeader(tc.upstream)
				if tc.body != "" {
					io.WriteString(w, tc.body)
					return
				}
				json.NewEncoder(w).Encode(map[string]any{"min": 0, "max": 4, "sum": 4, "average": 1, "anyDiagonal": false, "diagonalMatrices": []int{}, "matrixCount": tc.count})
			}))
			defer stub.Close()
			timeout := time.Second
			if tc.upstream == 0 {
				timeout = 100 * time.Millisecond
			}
			client, err := statistics.New(stub.URL, timeout)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest("POST", "/api/v1/matrices/analyze", strings.NewReader(`{"matrix":[[1,2],[3,4]],"operation":"`+tc.operation+`"}`))
			req.Header.Set("X-Request-ID", "test")
			res, err := New(client).Test(req, 2000)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != tc.status || res.Header.Get("X-Request-ID") != "test" {
				t.Fatalf("got %d", res.StatusCode)
			}
			if tc.count == 0 && calls.Load() != 0 {
				t.Fatal("invalid operation reached upstream")
			}
		})
	}
}
