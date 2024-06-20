package internal

import (
	"life-streams/internal/database"
	session_types "life-streams/internal/server/handlers/session/types"
)

type SessionQueries interface {
	GetSession(name string) (*session_types.Session, error)
	GetUserIDFromSession(sessionToken string) (int, error)
}

func GetSession(name string) (*session_types.Session, error) {
	database := database.New()

	var session session_types.Session

	query := `SELECT id, user_id, session_token, expires_at FROM sessions WHERE session_token = ?`
	row := database.QueryRow(query, name)
	err := row.Scan(&session.ID, &session.UserID, &session.SessionToken, &session.ExpiresAt)

	return &session, err
}

func GetUserIDFromSession(sessionToken string) (int, error) {
	database := database.New()

	var userID int
	query := `SELECT user_id FROM sessions WHERE session_token = ?`
	row := database.QueryRow(query, sessionToken)
	err := row.Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}
