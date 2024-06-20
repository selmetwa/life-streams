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

	fmt.Println("stream", stream)

	if err != nil {
		return stream, fmt.Errorf("something went wrong creating stream")
	}

	return stream, nil
}
