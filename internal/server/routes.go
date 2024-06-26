package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	web "life-streams/cmd/web"
	index "life-streams/cmd/web/pages/index"
	signup_view "life-streams/cmd/web/pages/signup"
	signup_handler "life-streams/internal/server/handlers/signup"

	login_view "life-streams/cmd/web/pages/login"
	login_handler "life-streams/internal/server/handlers/login"

	logout_handler "life-streams/internal/server/handlers/logout"
	session_handler "life-streams/internal/server/handlers/session"

	dashboard_handler "life-streams/internal/server/handlers/dashboard"
	stream_handler "life-streams/internal/server/handlers/stream"

	task_handler "life-streams/internal/server/handlers/tasks"

	"github.com/a-h/templ"
)

func authGatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("session_token")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		var token = session_handler.SessionHandler(sessionToken.Value)
		var sessionExpiresAt = token.ExpiresAt

		currentTime := time.Now()
		var isExpiredToken = sessionExpiresAt.Before(currentTime)

		if sessionToken.Value == "" || isExpiredToken {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func alreadyLoggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("session_token")

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		var token = session_handler.SessionHandler(sessionToken.Value)
		var sessionExpiresAt = token.ExpiresAt

		currentTime := time.Now()
		var isExpiredToken = sessionExpiresAt.Before(currentTime)

		if sessionToken.Value != "" && !isExpiredToken {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	mux.HandleFunc("/dashboard", dashboard_handler.DashboardHandler)
	mux.Handle("/", alreadyLoggedInMiddleware(templ.Handler(index.IndexPage(false))))

	mux.Handle("/signup", alreadyLoggedInMiddleware(templ.Handler(signup_view.SignupPage(false))))
	mux.HandleFunc("/signup_post", signup_handler.SignupHandler)

	mux.Handle("/login", alreadyLoggedInMiddleware(templ.Handler(login_view.LoginPage(false))))
	mux.HandleFunc("/login_post", login_handler.LoginHandler)

	mux.HandleFunc("/logout", logout_handler.LogoutHandler)

	mux.HandleFunc("/create_stream", stream_handler.CreateStream)
	mux.HandleFunc("/get_streams", stream_handler.RenderStreamList)

	mux.HandleFunc("/stream/{id}", stream_handler.StreamPage)
	mux.HandleFunc("/create_task", task_handler.CreateTask)
	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
