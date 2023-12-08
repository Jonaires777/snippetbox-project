package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.project/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	// httptest.NewRequest does not return an error, it only returns a *http.Request
	// so we don't need to check for an error
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
