package internal

import (
	"fmt"
	create_task_modal_view "life-streams/cmd/web/components/create_task_modal"
	task_list "life-streams/cmd/web/components/task_list"
	session_queries "life-streams/internal/server/handlers/session/queries"
	task_mutations "life-streams/internal/server/handlers/tasks/mutations"
	task_queries "life-streams/internal/server/handlers/tasks/queries"
	"net/http"
	"strconv"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	taskName := r.FormValue("title")
	description := r.FormValue("description")
	streamIdStr := r.FormValue("stream")
	streamId, _ := strconv.Atoi(streamIdStr)

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

	task, err := task_mutations.CreateTask(userId, streamId, taskName, description)

	if err != nil {
		component := create_task_modal_view.CreateTaskError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Task created successfully: ", task)
	w.Header().Set("HX-Trigger", "refetchStreamList")
	w.Header().Set("HX-Trigger", "refetchTasks")

	component := create_task_modal_view.CreateTaskSuccess()
	component.Render(r.Context(), w)
}

func RenderTaskList(w http.ResponseWriter, r *http.Request) {
	stream_id_str := r.PathValue("id")
	streamId, _ := strconv.Atoi(stream_id_str)

	tasks, err := task_queries.GetTaskByStreamID(streamId)

	if err != nil {
		fmt.Println("something went wrong getting tasks", err)
	}

	component := task_list.TaskList(tasks)
	component.Render(r.Context(), w)
}
