package stayease

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestDashboardHandler(t *testing.T) {
	server := NewServer()
	reqCleaner := httptest.NewRequest("GET", "/dashboard", nil)
	reqCleaner.AddCookie(&http.Cookie{Name: "session_user", Value: "cleaner"})
	reqCleaner.AddCookie(&http.Cookie{Name: "session_role", Value: "Cleaner"})
	wCleaner := httptest.NewRecorder()
	server.ServeHTTP(wCleaner, reqCleaner)

	body := wCleaner.Body.String()
	if !strings.Contains(body, "Room Inventory") {
		t.Errorf("cleaner should see Room Inventory")
	}
}

func TestCheckoutActionPermission(t *testing.T) {
	server := NewServer()

	// Cleaner checkout request should fail
	reqCleaner := httptest.NewRequest("POST", "/checkout?rate=150&days=3", nil)
	reqCleaner.AddCookie(&http.Cookie{Name: "session_role", Value: "Cleaner"})
	wCleaner := httptest.NewRecorder()
	server.ServeHTTP(wCleaner, reqCleaner)
	if wCleaner.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden for Cleaner checkout, got %d", wCleaner.Code)
	}

	// Receptionist checkout request should pass
	reqRecep := httptest.NewRequest("POST", "/checkout?rate=150&days=3", nil)
	reqRecep.AddCookie(&http.Cookie{Name: "session_role", Value: "Receptionist"})
	wRecep := httptest.NewRecorder()
	server.ServeHTTP(wRecep, reqRecep)
	if wRecep.Code != http.StatusOK {
		t.Errorf("expected 200 OK for Receptionist checkout, got %d", wRecep.Code)
	}
}
