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

	// Request dashboard without cookie should redirect to login
	reqUnauth := httptest.NewRequest("GET", "/dashboard", nil)
	wUnauth := httptest.NewRecorder()
	server.ServeHTTP(wUnauth, reqUnauth)
	if wUnauth.Code != http.StatusFound {
		t.Errorf("expected redirect 302, got %d", wUnauth.Code)
	}

	// Request dashboard with Cleaner cookie should only contain Cleaner actions
	reqCleaner := httptest.NewRequest("GET", "/dashboard", nil)
	reqCleaner.AddCookie(&http.Cookie{Name: "session_user", Value: "cleaner"})
	reqCleaner.AddCookie(&http.Cookie{Name: "session_role", Value: "Cleaner"})
	wCleaner := httptest.NewRecorder()
	server.ServeHTTP(wCleaner, reqCleaner)

	body := wCleaner.Body.String()
	if !strings.Contains(body, "Room Inventory") {
		t.Errorf("cleaner should see Room Inventory")
	}
	if strings.Contains(body, "Billing Processor") {
		t.Errorf("cleaner should NOT see Billing Processor")
	}
}
