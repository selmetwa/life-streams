package internal

import (
	"fmt"
	"life-streams/internal/database"
	auth_queries "life-streams/internal/server/handlers/auth/queries"
	auth_types "life-streams/internal/server/handlers/auth/types"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthQueries interface {
	AddUser(email, password string) error
	LoginUser(email, password string) (auth_types.User, string, error)
	LogoutUser(sessionToken string) error
}

func AddUser(email, password string) error {
	database := database.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
	_, err = database.Exec(query, email, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(email, password string) (*auth_types.User, string, error) {
	database := database.New()

	user, err := auth_queries.GetUserByEmail(email)

	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", fmt.Errorf("user not found with that email. please sign up")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("incorrect password. please try again")
	}

	// create session
	sessionToken := uuid.New().String()

	// insert session
	query := `INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)`
	_, err = database.Exec(query, user.ID, sessionToken, time.Now().Add(24*time.Hour))

	if err != nil {
		return nil, "", fmt.Errorf("error creating session for user")
	}

	return user, sessionToken, nil
}

func LogoutUser(sessionToken string) error {
	database := database.New()

	query := `DELETE FROM sessions WHERE session_token = ?`
	_, err := database.Exec(query, sessionToken)

	if err != nil {
		return err
	}

	return nil
}
