package stayease

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	server := NewServer()
	reqGet := httptest.NewRequest("GET", "/login", nil)
	wGet := httptest.NewRecorder()
	server.ServeHTTP(wGet, reqGet)
	if wGet.Code != http.StatusOK {
		t.Errorf("expected 200 OK for GET /login, got %d", wGet.Code)
	}
}

func TestLandingRedirect(t *testing.T) {
	server := NewServer()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code != http.StatusFound {
		t.Errorf("expected redirect from / to /login, got code %d", w.Code)
	}
}
