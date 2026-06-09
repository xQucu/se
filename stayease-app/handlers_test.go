package stayease

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	server := NewServer()

	// GET /login should serve login page
	reqGet := httptest.NewRequest("GET", "/login", nil)
	wGet := httptest.NewRecorder()
	server.ServeHTTP(wGet, reqGet)
	if wGet.Code != http.StatusOK {
		t.Errorf("expected 200 OK for GET /login, got %d", wGet.Code)
	}

	// POST /login with invalid pass
	reqPostBad := httptest.NewRequest("POST", "/login", strings.NewReader("username=owner&password=short"))
	reqPostBad.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	wPostBad := httptest.NewRecorder()
	server.ServeHTTP(wPostBad, reqPostBad)
	if wPostBad.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", wPostBad.Code)
	}
}
