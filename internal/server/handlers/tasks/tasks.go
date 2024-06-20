package internal

import (
	"fmt"
	create_task_modal_view "life-streams/cmd/web/components/create_task_modal"
	session_queries "life-streams/internal/server/handlers/session/queries"
	task_mutations "life-streams/internal/server/handlers/tasks/mutations"
	task_queries "life-streams/internal/server/handlers/tasks/queries"
	"net/http"
	"strconv"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	taskName := r.FormValue("title")
	description := r.FormValue("description")
	streamIdStr := r.FormValue("stream")
	streamId, _ := strconv.Atoi(streamIdStr)

	fmt.Println("Stream ID: ", streamId)

	sessionToken, _ := r.Cookie("session_token")

	userId, err := session_queries.GetUserIDFromSession(sessionToken.Value)

	if err != nil {
		component := create_task_modal_view.CreateTaskError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	task_id, _ := task_queries.GetTaskByTitle(userId, taskName)

	if task_id != nil {
		component := create_task_modal_view.CreateTaskError("Task with this name already exists")
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Task does not exist, creating task")

	task, err := task_mutations.CreateTask(userId, streamId, taskName, description)

	if err != nil {
		component := create_task_modal_view.CreateTaskError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Task created successfully: ", task)
	w.Header().Set("HX-Trigger", "refetchStreamList")

	component := create_task_modal_view.CreateTaskSuccess()
	component.Render(r.Context(), w)
}
