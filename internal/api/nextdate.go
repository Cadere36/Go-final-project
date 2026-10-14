package api

import (
	"net/http"
	"time"

	"go-final-project/internal/scheduler"
)

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeText(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		var err error
		now, err = time.Parse(scheduler.DateFormat, nowStr)
		if err != nil {
			writeText(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	next, err := scheduler.NextDate(now, date, repeat)
	if err != nil {
		writeText(w, http.StatusBadRequest, err.Error())
		return
	}
	writeText(w, http.StatusOK, next)
}
