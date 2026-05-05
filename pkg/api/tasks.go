package api

import (
	"net/http"

	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

type tasksResponse struct {
	Tasks []db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	// Получаем задачи из БД (ограничиваем 50 записями)
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка получения задач"}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, tasksResponse{Tasks: tasks}, http.StatusOK)
}
