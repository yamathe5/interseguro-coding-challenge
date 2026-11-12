package api

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInvalidRequestIDIsReplaced(t *testing.T) {
	for _, id := range []string{"unsafe value", strings.Repeat("x", 65)} {
		req := httptest.NewRequest("GET", "/health/live", nil)
		req.Header.Set("X-Request-ID", id)
		res, err := New(nil).Test(req)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		got := res.Header.Get("X-Request-ID")
		if got == id || !requestIDPattern.MatchString(got) {
			t.Errorf("invalid replacement ID: %q", got)
		}
	}
}
