package stayease

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//go:embed templates/login.html
var loginHTML []byte

//go:embed templates/dashboard.html
var dashboardTemplateRaw []byte

var mockRooms = []Room{
	{ID: "1", Number: "101", Status: "Available", Rate: 100},
	{ID: "2", Number: "102", Status: "Needs Cleaning", Rate: 120},
}

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

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})
	
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(loginHTML)
	})

	mux.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request) {
		// Clear cookies
		http.SetCookie(w, &http.Cookie{Name: "session_user", Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(w, &http.Cookie{Name: "session_role", Value: "", Path: "/", MaxAge: -1})
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if !ValidatePassword(password) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`<div class="text-red-500 text-sm">Password must contain at least 8 characters</div>`))
			return
		}

		user, ok := AuthenticateUser(username, password)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`<div class="text-red-500 text-sm">Invalid username or password</div>`))
			return
		}

		setSessionCookies(w, user.Username, user.Role)
		w.Header().Set("HX-Redirect", "/dashboard")
		w.Write([]byte("Login successful"))
	})

	mux.HandleFunc("GET /dashboard", func(w http.ResponseWriter, r *http.Request) {
		role, ok := getSessionRole(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		t, _ := template.New("dashboard").Parse(string(dashboardTemplateRaw))
		w.Header().Set("Content-Type", "text/html")
		_ = t.Execute(w, map[string]interface{}{
			"Role":             string(role),
			"CanCalculateBill": HasPermission(role, "calculate_bill"),
		})
	})

	mux.HandleFunc("GET /rooms", func(w http.ResponseWriter, r *http.Request) {
		role, ok := getSessionRole(r)
		if !ok || !HasPermission(role, "view_rooms") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		w.Header().Set("Content-Type", "text/html")
		for _, room := range mockRooms {
			fmt.Fprintf(w, `
				<div class="p-4 bg-white border shadow rounded flex justify-between items-center">
					<div>
						<span class="font-bold">Room %s</span> - <span class="text-gray-500 text-sm">%s</span>
					</div>
					<span class="font-semibold text-indigo-600">$%0.2f/night</span>
				</div>
			`, room.Number, room.Status, room.Rate)
		}
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
		w.Write([]byte(fmt.Sprintf(`<div class="text-xl font-bold text-green-600">Calculated Invoice: $%0.2f</div>`, bill)))
	})

	mux.HandleFunc("POST /status-update", func(w http.ResponseWriter, r *http.Request) {
		role, ok := getSessionRole(r)
		if !ok || !HasPermission(role, "update_room_status") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}

		_ = r.ParseForm()
		roomID := r.FormValue("room_id")
		status := r.FormValue("status")

		for i, room := range mockRooms {
			if room.ID == roomID {
				mockRooms[i].Status = status
				w.Write([]byte(fmt.Sprintf(`<div class="text-green-600 font-semibold">Room %s status updated to %s</div>`, room.Number, status)))
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Room not found"))
	})

	return mux
}
