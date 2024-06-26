package internal

import (
	"fmt"
	"life-streams/internal/database"
	stream_types "life-streams/internal/server/handlers/stream/types"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Stream struct {
	ID          int
	Title       string
	Description string
	Priority    int
	Position    int
	TasksCount  int
}

type StreamQueries interface {
	GetStreamsByUserID(userId int) ([]stream_types.Stream, error)
	GetStreamByTitle(user_id int, title string) (*int, error)
	GetStreamTitleById(user_id int, stream_id int) (*string, error)
}

func GetStreamsByUserID(userId int) ([]stream_types.Stream, error) {
	database := database.New()

	query := `SELECT id, title, description, priority, position FROM streams WHERE user_id = ? ORDER BY updated_at DESC`
	rows, err := database.Query(query, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var streams []stream_types.Stream

	for rows.Next() {
		var stream stream_types.Stream

		var count int

		err := rows.Scan(&stream.ID, &stream.Title, &stream.Description, &stream.Priority, &stream.Position)

		fmt.Println("stream", stream)
		fmt.Println("streamID", stream.ID)
		if err != nil {
			fmt.Println("error scanning stream", err)
			return nil, err
		}

		tasksCountQuery := `SELECT COUNT(*) FROM tasks WHERE stream_id = ? AND user_id = ?`
		database.QueryRow(tasksCountQuery, stream.ID, userId).Scan(&count)

		stream.TasksCount = count
		streams = append(streams, stream)
	}

	return streams, nil
}

func GetStreamByTitle(user_id int, title string) (*int, error) {
	database := database.New()

	query := `SELECT id from streams where title = ? AND user_id = ?`
	row := database.QueryRow(query, title, user_id)

	var id *int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func GetStreamTitleById(user_id int, stream_id int) (string, error) {
	database := database.New()
	query := `SELECT title from streams where id = ?`

	var title string

	err := database.QueryRow(query, stream_id).Scan(&title)
	if err != nil {
		fmt.Println("Error executing query:", err)
	}

	fmt.Println("title: ", title)

	return title, nil
}
