package api

import (
	"net/http"
	"time"

	"go-final-project/internal/db"
	"go-final-project/internal/scheduler"
)

const tasksLimit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	search := r.FormValue("search")

	var tasks []*db.Task
	var err error

	if search == "" {
		tasks, err = db.Tasks(tasksLimit)
	} else if t, perr := time.Parse("02.01.2006", search); perr == nil {
		tasks, err = db.TasksByDate(t.Format(scheduler.DateFormat), tasksLimit)
	} else {
		tasks, err = db.TasksBySearch(search, tasksLimit)
	}

	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, TasksResp{Tasks: tasks})
}
