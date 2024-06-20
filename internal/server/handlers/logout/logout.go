package web

import (
	"fmt"
	auth_mutations "life-streams/internal/server/handlers/auth/mutations"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")

	if err != nil {
		fmt.Println("error getting session token from cookie. Err: %v", err)
	}

	fmt.Println("sessionToken: ", sessionToken)

	err = auth_mutations.LogoutUser(sessionToken.Value)

	if err != nil {
		fmt.Println("error logging out user: ", err)
	}

	// Create a new cookie with the same name, but with an expired value
	expiredCookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1, // Set MaxAge to -1 to delete the cookie
	}

	fmt.Println("cookie: ", expiredCookie)
	http.SetCookie(w, expiredCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
