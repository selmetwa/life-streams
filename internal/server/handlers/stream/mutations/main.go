package internal

import (
	"fmt"
	"life-streams/internal/database"
	stream_types "life-streams/internal/server/handlers/stream/types"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type StreamMutation interface {
	CreateStream(userId int, title string, description string, priority int) (int, error)
	DeleteStream(userId int, streamID int)
}

type Stream struct {
	ID          int
	Title       string
	Description string
	Priority    int
	Position    int
	TasksCount  int
}

func CreateStream(user_id int, title, description string, priority int) (stream_types.Stream, error) {
	database := database.New()

	existing_streams_query := `SELECT position FROM streams WHERE user_id = ? ORDER BY position DESC LIMIT 1`

	row := database.QueryRow(existing_streams_query, user_id)
	var position int
	err := row.Scan(&position)

	if err != nil {
		position = 0
	}

	var stream stream_types.Stream
	query := `INSERT INTO streams (user_id, title, description, priority, position) VALUES (?, ?, ?, ?, ?)`
	res, _ := database.Exec(query, user_id, title, description, priority, position+1)
	lastInsertId, _ := res.LastInsertId()

	query = `SELECT id, title, description, priority, position FROM streams WHERE id = ? AND user_id = ?`
	stream_row := database.QueryRow(query, lastInsertId, user_id)

	// Step 4: Assign the row values to your stream variable
	err = stream_row.Scan(&stream.ID, &stream.Title, &stream.Description, &stream.Priority, &stream.Position)

	if err != nil {
		return stream, err // handle error
	}

	if err != nil {
		return stream, fmt.Errorf("something went wrong creating stream")
	}

	return stream, nil
}

func DeleteStream(userId int, streamID int) error {
	database := database.New()

	delete_tasks_from_stream := `DELETE FROM tasks WHERE user_id = ? AND stream_id = ?`

	_, err := database.Exec(delete_tasks_from_stream, userId, streamID)

	if err != nil {
		fmt.Println("something went wrong deleting tasks within stream: ", err)
		return err
	}

	delete_stream_mutation := `DELETE FROM streams WHERE user_id = ? AND id = ?`
	row, err := database.Exec(delete_stream_mutation, userId, streamID)

	if err != nil {
		fmt.Println("something went wrong deleting stream: ", err)
		return err
	}

	fmt.Println("row: ", row)

	return nil
}
