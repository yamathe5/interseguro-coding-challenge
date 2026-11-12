package api

import (
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	for _, path := range []string{"/health/live", "/health/ready"} {
		res, err := New().Test(httptest.NewRequest("GET", path, nil))
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		if res.StatusCode != 200 || res.Header.Get("X-Request-ID") == "" {
			t.Fatalf("%s: %v", path, res)
		}
	}
}
