package internal

import (
	"fmt"
	dashboard_view "life-streams/cmd/web/pages/dashboard"
	session_handler "life-streams/internal/server/handlers/session"
	session_queries "life-streams/internal/server/handlers/session/queries"
	stream_queries "life-streams/internal/server/handlers/stream/queries"
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

	userId, _ := session_queries.GetUserIDFromSession(sessionToken.Value)

	streams, _ := stream_queries.GetStreamsByUserID(userId)

	fmt.Println(streams)

	component := dashboard_view.Dashboard(true, streams)
	component.Render(r.Context(), w)
}
