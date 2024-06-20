package internal

import (
	"fmt"
	"life-streams/internal/database"
	task_types "life-streams/internal/server/handlers/tasks/types"
)

type TaskMutations interface {
}

func CreateTask(user_id int, stream_id int, title string, description string) (task_types.Task, error) {
	database := database.New()
	var Task task_types.Task

	var create_task_statement = `INSERT into tasks (user_id, stream_id, title, description, priority, position) VALUES (?,?,?,?,?,?)`

	res, err := database.Exec(create_task_statement, user_id, stream_id, title, description, 0, 0)

	if err != nil {
		fmt.Println("error creating task", err)
	}

	lastInsertId, _ := res.LastInsertId()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("lastInsertId", lastInsertId)

	query := `SELECT id, title, description, stream_id FROM tasks WHERE id = ? AND user_id = ?`
	task_row := database.QueryRow(query, lastInsertId, user_id)

	fmt.Println("task_row", task_row)
	// Step 4: Assign the row values to your stream variable
	err = task_row.Scan(&Task.ID, &Task.Title, &Task.Description, &Task.StreamID)

	if err != nil {
		fmt.Println("error", err)
	}

	return Task, nil
}
