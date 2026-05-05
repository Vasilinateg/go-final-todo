package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
}

// parseDate парсит дату в формате YYYYMMDD и возвращает UTC полночь
func parseDate(s string) (time.Time, error) {
	if len(s) != 8 {
		return time.Time{}, strconv.ErrSyntax
	}
	year, _ := strconv.Atoi(s[0:4])
	month, _ := strconv.Atoi(s[4:6])
	day, _ := strconv.Atoi(s[6:8])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	nowStr := strings.TrimSpace(r.URL.Query().Get("now"))
	dateStr := strings.TrimSpace(r.URL.Query().Get("date"))
	repeat := strings.TrimSpace(r.URL.Query().Get("repeat"))

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now().UTC()
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	} else {
		now, err = parseDate(nowStr)
		if err != nil {
			http.Error(w, "bad now", http.StatusBadRequest)
			return
		}
	}

	if dateStr == "" || repeat == "" {
		http.Error(w, "missing date or repeat", http.StatusBadRequest)
		return
	}

	// Валидируем date
	if _, err := parseDate(dateStr); err != nil {
		http.Error(w, "bad date", http.StatusBadRequest)
		return
	}

	result, err := NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
