package main

import (
	"fmt"
	"life-streams/internal/server"
)

func main() {
	fmt.Println("Starting server")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
