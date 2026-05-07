package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
			now, err := time.Parse(DateFormat, tt.nowStr)
			assert.NoError(t, err)
			result, err := NextDate(now, tt.dateStr, tt.repeat)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
