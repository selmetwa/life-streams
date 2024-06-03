package server

import (
	"encoding/json"
	"log"
	"net/http"

	hello "life-streams/cmd/web/components/hello"
	index "life-streams/cmd/web/components/index"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(index.IndexPage()))
	mux.HandleFunc("/health", s.healthHandler)
	mux.Handle("/web", templ.Handler(hello.HelloForm()))
	mux.HandleFunc("/hello", hello.HelloWebHandler)

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
