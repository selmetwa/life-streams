package internal

import (
	"fmt"
	"life-streams/internal/database"
	task_types "life-streams/internal/server/handlers/tasks/types"
)

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

func GetTaskByStreamID(userId int, streamId int) ([]task_types.Task, error) {
	database := database.New()
	tasksQuery := `SELECT id, stream_id, title, description FROM tasks WHERE user_id = ? AND stream_id = ? ORDER BY updated_at DESC`

	rows, err := database.Query(tasksQuery, userId, streamId)

	fmt.Println("task_rows: ", rows)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []task_types.Task

	for rows.Next() {
		var task task_types.Task
		err := rows.Scan(&task.ID, &task.StreamID, &task.Title, &task.Description)

		if err != nil {
			fmt.Println("something went wrong converting rows to arr", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
