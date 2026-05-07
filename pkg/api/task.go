package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		writeJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
	}
}

// POST /api/task
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	now := time.Now()
	nowDate := now.Format(DateFormat)

	if task.Date == "" {
		task.Date = nowDate
	}

	dateParsed, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Неверный формат даты"}, http.StatusBadRequest)
		return
	}

	nowNormalized := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dateNormalized := time.Date(dateParsed.Year(), dateParsed.Month(), dateParsed.Day(), 0, 0, 0, 0, dateParsed.Location())

	if dateNormalized.Before(nowNormalized) {
		if task.Repeat == "" {
			task.Date = nowDate
		} else {
			nextDate, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
				return
			}
			task.Date = nextDate
		}
	}

	id, err := db.AddTask(task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка добавления задачи"}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{"id": id}, http.StatusOK)
}

// GET /api/task?id=123
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Задача не найдена"}, http.StatusNotFound)
		return
	}

	writeJSON(w, task, http.StatusOK)
}

// PUT /api/task
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор задачи"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	now := time.Now()
	nowDate := now.Format(DateFormat)

	if task.Date == "" {
		task.Date = nowDate
	}

	dateParsed, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Неверный формат даты"}, http.StatusBadRequest)
		return
	}

	nowNormalized := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dateNormalized := time.Date(dateParsed.Year(), dateParsed.Month(), dateParsed.Day(), 0, 0, 0, 0, dateParsed.Location())

	if dateNormalized.Before(nowNormalized) {
		if task.Repeat == "" {
			task.Date = nowDate
		} else {
			nextDate, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
				return
			}
			task.Date = nextDate
		}
	}

	if err := db.UpdateTask(task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	writeJSON(w, map[string]interface{}{}, http.StatusOK)
}

// DELETE /api/task?id=123
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"}, http.StatusBadRequest)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	writeJSON(w, map[string]interface{}{}, http.StatusOK)
}
