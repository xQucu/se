package stayease

import (
	"net/http"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<form hx-post="/login" hx-target="#login-err" class="space-y-4">
				<input type="text" name="username" placeholder="Username" required>
				<input type="password" name="password" placeholder="Password" required>
				<button type="submit">Login</button>
				<div id="login-err"></div>
			</form>
		`))
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if !ValidatePassword(password) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Password must contain at least 8 characters"))
			return
		}

		user, ok := AuthenticateUser(username, password)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid username or password"))
			return
		}

		// Success session cookie mock
		http.SetCookie(w, &http.Cookie{
			Name:  "session_user",
			Value: string(user.Username),
			Path:  "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "session_role",
			Value: string(user.Role),
			Path:  "/",
		})
		w.Header().Set("HX-Redirect", "/dashboard")
		w.Write([]byte("Login successful"))
	})
	return mux
}
