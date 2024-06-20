package internal

import (
	"database/sql"
	"fmt"
	"life-streams/internal/database"
	auth_types "life-streams/internal/server/handlers/auth/types"
)

func GetUserByEmail(email string) (*auth_types.User, error) {
	var user auth_types.User
	database := database.New()

	query := `SELECT id, username, password_hash, email, created_at, updated_at FROM users WHERE email = ?`
	row := database.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Password_hash, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given email
		}
		return nil, fmt.Errorf("error querying user by email")
	}

	return &user, nil
}
