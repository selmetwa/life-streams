package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Stream struct {
	ID          int
	Title       string
	Description string
	Priority    int
	Position    int
	TasksCount  int
}

type Task struct {
	ID          int
	Title       string
	Description string
	StreamID    int
}

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	GetUserByEmail(email string) (*User, error)
	AddUser(email, password string) error
	LoginUser(email, password string) (*User, string, error)
	LogoutUser(sessionToken string) error
	GetSession(name string) (*Session, error)
	GetUserIDFromSession(sessionToken string) (int, error)
	GetStreamByTitle(user_id int, title string) (*int, error)
	CreateStream(user_id int, title, description string, priority int) (Stream, error)
	GetStreamsByUserID(userId int) ([]Stream, error)
	GetTaskByTitle(userId int, title string) (*int, error)
	CreateTask(user_id int, stream_id int, title string, description string) (Task, error)
}

type service struct {
	db *sql.DB
}

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

func CreateTables(db *sql.DB) {
	createUsersStatement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		fmt.Println("create users table error: ", err)
	}

	createUsersStatement.Exec()

	createSessionsStatement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			session_token TEXT UNIQUE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			expires_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`)

	if err != nil {
		fmt.Println("create sessions table error: ", err)
	}

	createSessionsStatement.Exec()

	createStreamsStatement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS streams (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			priority INTEGER NOT NULL,
			position INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)

	if err != nil {
		fmt.Println("create streams table error: ", err)
	}

	createStreamsStatement.Exec()

	createTasksStatement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			stream_id INTEGER NOT NULL,
			priority INTEGER NOT NULL,
			title TEXT NOT NULL,
			description TEXT,
			due_date DATETIME,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			position INTEGER NOT NULL,
			FOREIGN KEY (stream_id) REFERENCES streams(id) ON DELETE CASCADE
		)
	`)

	if err != nil {
		fmt.Println("create streams tasks error: ", err)
	}

	createTasksStatement.Exec()
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}

	CreateTables(db)
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}

type User struct {
	id            int
	username      string
	password_hash string
	email         string
	created_at    string
	updated_at    string
}

func (s *service) GetUserByEmail(email string) (*User, error) {
	var user User

	query := `SELECT id, username, password_hash, email, created_at, updated_at FROM users WHERE email = ?`
	row := s.db.QueryRow(query, email)
	err := row.Scan(&user.id, &user.username, &user.password_hash, &user.email, &user.created_at, &user.updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given email
		}
		return nil, fmt.Errorf("error querying user by email")
	}

	return &user, nil
}

func (s *service) AddUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
	_, err = s.db.Exec(query, email, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) LoginUser(email, password string) (*User, string, error) {
	user, err := s.GetUserByEmail(email)

	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", fmt.Errorf("user not found with that email. please sign up")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.password_hash), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("incorrect password. please try again")
	}

	// create session
	sessionToken := uuid.New().String()

	// insert session
	query := `INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)`
	_, err = s.db.Exec(query, user.id, sessionToken, time.Now().Add(24*time.Hour))

	if err != nil {
		return nil, "", fmt.Errorf("error creating session for user")
	}

	return user, sessionToken, nil
}

func (s *service) LogoutUser(sessionToken string) error {
	query := `DELETE FROM sessions WHERE session_token = ?`
	_, err := s.db.Exec(query, sessionToken)

	if err != nil {
		return err
	}

	return nil
}

type Session struct {
	ID           int
	UserID       int
	SessionToken string
	ExpiresAt    time.Time
}

func (s *service) GetSession(name string) (*Session, error) {
	var session Session

	query := `SELECT id, user_id, session_token, expires_at FROM sessions WHERE session_token = ?`
	row := s.db.QueryRow(query, name)
	err := row.Scan(&session.ID, &session.UserID, &session.SessionToken, &session.ExpiresAt)

	return &session, err
}

func (s *service) GetUserIDFromSession(sessionToken string) (int, error) {
	var userID int
	fmt.Println("sessionToken", sessionToken)
	query := `SELECT user_id FROM sessions WHERE session_token = ?`
	row := s.db.QueryRow(query, sessionToken)
	err := row.Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *service) GetStreamByTitle(user_id int, title string) (*int, error) {
	query := `SELECT id from streams where title = ? AND user_id = ?`
	row := s.db.QueryRow(query, title, user_id)

	var id *int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *service) CreateStream(user_id int, title, description string, priority int) (Stream, error) {
	existing_streams_query := `SELECT position FROM streams WHERE user_id = ? ORDER BY position DESC LIMIT 1`

	row := s.db.QueryRow(existing_streams_query, user_id)
	var position int
	err := row.Scan(&position)

	if err != nil {
		position = 0
	}

	var stream Stream
	query := `INSERT INTO streams (user_id, title, description, priority, position) VALUES (?, ?, ?, ?, ?)`
	res, _ := s.db.Exec(query, user_id, title, description, priority, position+1)
	lastInsertId, _ := res.LastInsertId()

	query = `SELECT id, title, description, priority, position FROM streams WHERE id = ? AND user_id = ?`
	stream_row := s.db.QueryRow(query, lastInsertId, user_id)

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

func (s *service) GetStreamsByUserID(userId int) ([]Stream, error) {
	query := `SELECT id, title, description, priority, position FROM streams WHERE user_id = ? ORDER BY updated_at DESC`
	rows, err := s.db.Query(query, userId)

	tasksCountQuery := `SELECT COUNT(*) FROM tasks WHERE stream_id = ? AND user_id = ?`

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var streams []Stream
	for rows.Next() {
		var stream Stream

		// countRow, _ := s.db.Query(tasksCountQuery, stream.ID, userId)

		var count int

		s.db.QueryRow(tasksCountQuery, stream.ID, userId).Scan(&count)

		fmt.Println("count", count)

		err := rows.Scan(&stream.ID, &stream.Title, &stream.Description, &stream.Priority, &stream.Position)

		if err != nil {
			return nil, err
		}

		stream.TasksCount = count
		streams = append(streams, stream)
	}

	return streams, nil
}

func (s *service) GetTaskByTitle(userId int, title string) (*int, error) {
	taskQuery := `SELECT id from tasks where user_id = ? AND title = ?`
	row := s.db.QueryRow(taskQuery, userId, title)

	var id *int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *service) CreateTask(user_id int, stream_id int, title string, description string) (Task, error) {
	var Task Task

	var create_task_statement = `INSERT into tasks (user_id, stream_id, title, description, priority, position) VALUES (?,?,?,?,?,?)`

	res, err := s.db.Exec(create_task_statement, user_id, stream_id, title, description, 0, 0)

	if err != nil {
		fmt.Println("error creating task", err)
	}

	lastInsertId, _ := res.LastInsertId()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("lastInsertId", lastInsertId)

	query := `SELECT id, title, description, stream_id FROM tasks WHERE id = ? AND user_id = ?`
	task_row := s.db.QueryRow(query, lastInsertId, user_id)

	fmt.Println("task_row", task_row)
	// Step 4: Assign the row values to your stream variable
	err = task_row.Scan(&Task.ID, &Task.Title, &Task.Description, &Task.StreamID)

	if err != nil {
		fmt.Println("error", err)
	}

	return Task, nil
}
