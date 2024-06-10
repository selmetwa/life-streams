package internal

import (
	"fmt"
	signup_view "life-streams/cmd/web/components/signup"
	db "life-streams/internal/database"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		component := signup_view.SignUpError("Something went wrong parsing form. Please try again.")
		_ = component.Render(r.Context(), w)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var instance = db.New()
	// check if email already exists
	user, _ := instance.GetUserByEmail(email)

	// if user is nil, create user
	if user == nil {
		err := instance.AddUser(email, password)

		if err != nil {
			component := signup_view.SignUpError("Something went wrong creating user")
			component.Render(r.Context(), w)
		} else {
			component := signup_view.SignUpSuccess()
			component.Render(r.Context(), w)
		}
	} else {
		component := signup_view.SignUpError("User already exists with this email")
		component.Render(r.Context(), w)
	}
}
