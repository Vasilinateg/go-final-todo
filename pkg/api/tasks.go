package api

import (
	"net/http"

	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

const tasksLimit = 50

type tasksResponse struct {
	Tasks []db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	tasks, err := db.Tasks(tasksLimit)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка получения задач"}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, tasksResponse{Tasks: tasks}, http.StatusOK)
}
