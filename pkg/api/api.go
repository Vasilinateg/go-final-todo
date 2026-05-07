package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const DateFormat = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", taskDoneHandler)
}

// writeJSON отправляет JSON-ответ
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
	}
}

// nextDateHandler обрабатывает GET-запросы /api/nextdate
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	nowStr := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(DateFormat, nowStr)
		if err != nil {
			writeJSON(w, map[string]string{"error": "bad now"}, http.StatusBadRequest)
			return
		}
	}

	if date == "" || repeat == "" {
		writeJSON(w, map[string]string{"error": "missing date or repeat"}, http.StatusBadRequest)
		return
	}

	result, err := NextDate(now, date, repeat)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(result)); err != nil {
		http.Error(w, "Ошибка записи ответа", http.StatusInternalServerError)
	}
}
