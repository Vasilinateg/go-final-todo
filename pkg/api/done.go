package api

import (
	"net/http"
	"time"

	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

// taskDoneHandler обрабатывает POST /api/task/done
func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"}, http.StatusBadRequest)
		return
	}

	// Получаем задачу
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Задача не найдена"}, http.StatusNotFound)
		return
	}

	// Если нет повторения — удаляем задачу
	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			return
		}
		writeJSON(w, map[string]interface{}{}, http.StatusOK)
		return
	}

	// Если есть повторение — вычисляем следующую дату
	now := time.Now()
	nextDate, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Обновляем дату задачи
	if err := db.UpdateTaskDate(id, nextDate); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{}, http.StatusOK)
}
