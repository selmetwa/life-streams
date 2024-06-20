package internal

import "life-streams/internal/database"

type TaskQueries interface {
	GetTaskByTitle(userId int, title string) (*int, error)
}

func GetTaskByTitle(userId int, title string) (*int, error) {
	database := database.New()
	taskQuery := `SELECT id from tasks where user_id = ? AND title = ?`
	row := database.QueryRow(taskQuery, userId, title)

	var id *int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}
