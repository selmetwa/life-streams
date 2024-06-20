package internal

import (
	session_queries "life-streams/internal/server/handlers/session/queries"
	session_types "life-streams/internal/server/handlers/session/types"
)

func SessionHandler(name string) *session_types.Session {
	session, _ := session_queries.GetSession(name)
	return session
}
