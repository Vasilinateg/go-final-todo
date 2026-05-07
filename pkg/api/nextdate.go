package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func normalize(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// NextDate вычисляет следующую дату выполнения задачи
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	now = normalize(now)
	date = normalize(date)

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("неправильный формат правила d")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil || days < 1 || days > 400 {
			return "", fmt.Errorf("некорректный интервал для d: от 1 до 400")
		}
		next := date
		next = next.AddDate(0, 0, days)
		for next.Before(now) || next.Equal(now) {
			next = next.AddDate(0, 0, days)
		}
		return next.Format(DateFormat), nil

	case "y":
		if len(parts) != 1 {
			return "", fmt.Errorf("неправильный формат правила y")
		}
		next := date
		next = next.AddDate(1, 0, 0)
		for next.Before(now) || next.Equal(now) {
			next = next.AddDate(1, 0, 0)
		}
		if next.Month() == 2 && next.Day() == 29 && !isLeapYear(next.Year()) {
			next = time.Date(next.Year(), 3, 1, 0, 0, 0, 0, next.Location())
		}
		return next.Format(DateFormat), nil

	default:
		return "", fmt.Errorf("неподдерживаемый формат правила: %s", rule)
	}
}
