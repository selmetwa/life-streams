package web

import (
	"fmt"
	login_view "life-streams/cmd/web/components/login"
	auth_mutations "life-streams/internal/server/handlers/auth/mutations"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// check if email already exists
	user, sessionToken, err := auth_mutations.LoginUser(email, password)

	if err != nil {
		fmt.Println("Error logging in user:", err)
		component := login_view.LoginFailure(err.Error())
		component.Render(r.Context(), w)
	}

	if user != nil && sessionToken != "" {
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			HttpOnly: true,
			MaxAge:   86400,
			Expires:  time.Now().Add(24 * time.Hour),
		}

		http.SetCookie(w, cookie)
		w.Header().Set("HX-Redirect", "/dashboard")
	}
}
