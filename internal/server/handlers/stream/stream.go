package internal

import (
	"fmt"
	create_stream_view "life-streams/cmd/web/components/create_stream_modal"
	streams_list_view "life-streams/cmd/web/components/streams_list"
	stream_page "life-streams/cmd/web/pages/stream"
	session_handler "life-streams/internal/server/handlers/session"
	session_queries "life-streams/internal/server/handlers/session/queries"
	stream_mutations "life-streams/internal/server/handlers/stream/mutations"
	stream_queries "life-streams/internal/server/handlers/stream/queries"
	task_queries "life-streams/internal/server/handlers/tasks/queries"
	"net/http"
	"strconv"
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

	userId, err := session_queries.GetUserIDFromSession(sessionToken.Value)

	if err != nil {
		component := create_stream_view.CreateStreamError(err.Error())
		component.Render(r.Context(), w)
		return

	}

	stream_id, _ := stream_queries.GetStreamByTitle(userId, streamName)

	if stream_id != nil {
		component := create_stream_view.CreateStreamError("Stream with this name already exists")
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Stream does not exist, creating stream")

	stream, err := stream_mutations.CreateStream(userId, streamName, description, priorityAsInt)

	if err != nil {
		component := create_stream_view.CreateStreamError(err.Error())
		component.Render(r.Context(), w)
	}

	w.Header().Set("HX-Trigger", "refetchStreamList")

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

	userId, _ := session_queries.GetUserIDFromSession(sessionToken.Value)

	streams, _ := stream_queries.GetStreamsByUserID(userId)

	fmt.Println(streams)

	component := streams_list_view.StreamsList(streams)
	component.Render(r.Context(), w)
}

func StreamPage(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")

	stream_id_str := r.PathValue("id")
	streamId, _ := strconv.Atoi(stream_id_str)

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

	userId, _ := session_queries.GetUserIDFromSession(sessionToken.Value)

	tasks, err := task_queries.GetTaskByStreamID(streamId)
	title, _ := stream_queries.GetStreamTitleById(userId, streamId)

	streams, _ := stream_queries.GetStreamsByUserID(userId)

	if err != nil {
		fmt.Println("something went wrong getting tasks", err)
	}

	component := stream_page.StreamPage(true, tasks, title, stream_id_str, streams)
	component.Render(r.Context(), w)
}

func DeleteStream(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE STREAM")
	sessionToken, err := r.Cookie("session_token")

	streamID := r.FormValue("streamID")
	streamIDInt, _ := strconv.Atoi(streamID)

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

	userId, _ := session_queries.GetUserIDFromSession(sessionToken.Value)

	err = stream_mutations.DeleteStream(userId, streamIDInt)

	if err != nil {
		component := stream_page.DeleteStreamError()
		component.Render(r.Context(), w)
	}

	component := stream_page.DeleteStreamSuccess()
	component.Render(r.Context(), w)
}
