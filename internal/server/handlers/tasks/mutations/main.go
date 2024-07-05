package internal

import (
	"fmt"
	"life-streams/internal/database"
	task_types "life-streams/internal/server/handlers/tasks/types"
)

type TaskMutations interface {
	CreateTask(user_id int, stream_id int, title string, description string) (task_types.Task, error)
	DeleteTask(taskID int) error
	EditTask(taskID int, taskName string, description string, streamID int) error
}

func DeleteTask(taskID int) error {
	database := database.New()
	var delete_task_statement = `DELETE FROM tasks where id = ?`

	_, err := database.Exec(delete_task_statement, taskID)

	if err != nil {
		fmt.Println("something went wrong deleting task: ", err)
	}

	return nil
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

	query := `SELECT id, title, description, stream_id FROM tasks WHERE id = ? AND user_id = ?`
	task_row := database.QueryRow(query, lastInsertId, user_id)

	err = task_row.Scan(&Task.ID, &Task.Title, &Task.Description, &Task.StreamID)

	if err != nil {
		fmt.Println("error", err)
	}

	return Task, nil
}

func EditTask(taskID int, taskName string, description string, streamID int) error {
	database := database.New()

	var edit_task_statement = `UPDATE tasks SET title = ?, description = ?, stream_id = ? WHERE id = ?`
	res, err := database.Exec(edit_task_statement, taskName, description, streamID, taskID)

	if err != nil {
		fmt.Println("something went wrong updating task: ", err)

		return err
	}

	fmt.Println("res: ", res)
	return nil
}
