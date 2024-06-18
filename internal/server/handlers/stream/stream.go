package internal

import (
	"fmt"
	create_stream_view "life-streams/cmd/web/components/create_stream_modal"
	streams_list_view "life-streams/cmd/web/components/streams_list"
	db "life-streams/internal/database"
	session_handler "life-streams/internal/server/handlers/session"
	"net/http"
	"time"
)

func CreateStream(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	streamName := r.FormValue("title")
	description := r.FormValue("description")
	priority := r.FormValue("priority")

	priority_map := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	var priorityAsInt int
	var ok bool

	if priorityAsInt, ok = priority_map[priority]; !ok {
		priorityAsInt = 1 // Fallback to 1 if the key is not found in the map
	}

	sessionToken, _ := r.Cookie("session_token")

	var instance = db.New()
	userId, err := instance.GetUserIDFromSession(sessionToken.Value)

	if err != nil {
		component := create_stream_view.CreateStreamError(err.Error())
		component.Render(r.Context(), w)
		return

	}

	stream_id, _ := instance.GetStreamByTitle(userId, streamName)

	if stream_id != nil {
		component := create_stream_view.CreateStreamError("Stream with this name already exists")
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Stream does not exist, creating stream")

	stream, err := instance.CreateStream(userId, streamName, description, priorityAsInt)

	if err != nil {
		component := create_stream_view.CreateStreamError(err.Error())
		component.Render(r.Context(), w)
	}

	w.Header().Set("HX-Trigger", "newStream")

	component := create_stream_view.CreateStreamSuccess(stream)
	component.Render(r.Context(), w)
}

func RenderStreamList(w http.ResponseWriter, r *http.Request) {
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

	var instance = db.New()
	userId, _ := instance.GetUserIDFromSession(sessionToken.Value)

	streams, _ := instance.GetStreamsByUserID(userId)

	fmt.Println(streams)

	component := streams_list_view.StreamsList(streams)
	component.Render(r.Context(), w)
}
