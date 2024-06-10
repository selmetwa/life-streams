package internal

import (
	"fmt"
	db "life-streams/internal/database"
)

func SessionHandler(name string) *db.Session {
	var instance = db.New()
	session, _ := instance.GetSession(name)
	fmt.Println("session 2", session.ExpiresAt)
	return session
}
