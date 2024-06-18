package internal

import (
	"fmt"
	dashboard_view "life-streams/cmd/web/components/dashboard"
	db "life-streams/internal/database"
	session_handler "life-streams/internal/server/handlers/session"
	"net/http"
	"time"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var token = session_handler.SessionHandler(sessionToken.Value)
	var sessionExpiresAt = token.ExpiresAt

	currentTime := time.Now()
	var isExpiredToken = sessionExpiresAt.Before(currentTime)

	if sessionToken.Value == "" || isExpiredToken {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var instance = db.New()
	userId, _ := instance.GetUserIDFromSession(sessionToken.Value)

	streams, _ := instance.GetStreamsByUserID(userId)

	fmt.Println(streams)

	component := dashboard_view.Dashboard(true, streams)
	component.Render(r.Context(), w)
}
