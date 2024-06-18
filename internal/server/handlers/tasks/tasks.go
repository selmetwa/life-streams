package internal

import (
	"fmt"
	create_task_modal_view "life-streams/cmd/web/components/create_task_modal"
	db "life-streams/internal/database"
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

	var instance = db.New()
	userId, err := instance.GetUserIDFromSession(sessionToken.Value)

	if err != nil {
		component := create_task_modal_view.CreateTaskError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	task_id, _ := instance.GetTaskByTitle(userId, taskName)

	if task_id != nil {
		component := create_task_modal_view.CreateTaskError("Task with this name already exists")
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Task does not exist, creating task")

	task, err := instance.CreateTask(userId, streamId, taskName, description)

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
