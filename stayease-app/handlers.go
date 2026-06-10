package stayease

import (
	"fmt"
	"net/http"
	"strconv"
)

func setSessionCookies(w http.ResponseWriter, username string, role Role) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_user",
		Value: username,
		Path:  "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "session_role",
		Value: string(role),
		Path:  "/",
	})
}

func getSessionRole(r *http.Request) (Role, bool) {
	cookie, err := r.Cookie("session_role")
	if err != nil {
		return "", false
	}
	return Role(cookie.Value), true
}

func writeRoomsSection(w http.ResponseWriter, role Role) {
	fmt.Fprintf(w, "<div>Room Inventory</div>")
}

func writeBillingSection(w http.ResponseWriter, role Role) {
	if HasPermission(role, "calculate_bill") {
		fmt.Fprintf(w, "<div>Billing Processor</div>")
	}
}

func parseQueryParamFloat(r *http.Request, key string) float64 {
	val := r.FormValue(key)
	if val == "" {
		val = r.URL.Query().Get(key)
	}
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func parseQueryParamInt(r *http.Request, key string) int {
	val := r.FormValue(key)
	if val == "" {
		val = r.URL.Query().Get(key)
	}
	i, _ := strconv.Atoi(val)
	return i
}

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form hx-post="/login">Login Form</form>`))
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		user, ok := AuthenticateUser(r.FormValue("username"), r.FormValue("password"))
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid login"))
			return
		}
		setSessionCookies(w, user.Username, user.Role)
		w.Header().Set("HX-Redirect", "/dashboard")
	})

	mux.HandleFunc("GET /dashboard", func(w http.ResponseWriter, r *http.Request) {
		role, ok := getSessionRole(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Dashboard</h1>")
		writeRoomsSection(w, role)
		writeBillingSection(w, role)
	})

	mux.HandleFunc("POST /checkout", func(w http.ResponseWriter, r *http.Request) {
		role, ok := getSessionRole(r)
		if !ok || !HasPermission(role, "calculate_bill") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}

		_ = r.ParseForm()
		rate := parseQueryParamFloat(r, "rate")
		days := parseQueryParamInt(r, "days")
		
		bill := CalculateBill(rate, days)
		w.Write([]byte(fmt.Sprintf("Total: $%0.2f", bill)))
	})
	return mux
}
