package web

import (
	"fmt"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignupHandler")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	component := SignupResponse(email, password)
	err = component.Render(r.Context(), w)
	fmt.Println("email:", email)
	fmt.Println("password:", password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error rendering in HelloWebHandler:", err)
	}
}
