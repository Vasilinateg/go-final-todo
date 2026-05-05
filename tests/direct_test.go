package tests

import (
	"testing"
	"time"

	"github.com/Vasilinateg/go-final-todo/pkg/api"
)

func TestNextDateDirect(t *testing.T) {
	tests := []struct {
		nowStr   string
		dateStr  string
		repeat   string
		expected string
	}{
		{"20250701", "20250701", "y", "20260701"},
		{"20240229", "20240229", "y", "20250301"},
		{"20240301", "20240301", "y", "20250301"},
		{"20240202", "20240202", "d 30", "20240303"},
		{"20240228", "20240228", "d 1", "20240229"},
	}

	for _, tt := range tests {
		t.Run(tt.repeat, func(t *testing.T) {
			now, err := time.Parse(api.DateFormat, tt.nowStr)
			if err != nil {
				t.Fatal(err)
			}
			result, err := api.NextDate(now, tt.dateStr, tt.repeat)
			if err != nil {
				t.Errorf("Ошибка: %v", err)
			}
			if result != tt.expected {
				t.Errorf("NextDate(%s, %s, %s) = %s, ожидается %s", tt.nowStr, tt.dateStr, tt.repeat, result, tt.expected)
			}
		})
	}
}
