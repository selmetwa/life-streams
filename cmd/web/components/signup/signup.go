package web

import (
	"fmt"
	db "life-streams/internal/database"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)

		component := SignupErrorResponse()
		_ = component.Render(r.Context(), w)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// check if email already exists
	var instance = db.New()
	user, _ := instance.GetUserByEmail(email)

	// if user is nil, create user
	if user == nil {
		err := instance.AddUser(email, password)

		if err != nil {
			component := SignupErrorResponseWithMessage("Something went wrong creating user", true)
			component.Render(r.Context(), w)
		} else {
			component := SignupErrorResponseWithMessage("User created successfully", false)
			component.Render(r.Context(), w)
		}
	} else {
		component := SignupErrorResponseWithMessage("User already exists with this email", true)
		err = component.Render(r.Context(), w)
		fmt.Println("User already exists")
	}
}
