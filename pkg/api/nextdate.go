package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	// Нормализуем время
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

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
		// Всегда добавляем интервал
		next = next.AddDate(0, 0, days)
		// Повторяем, пока не станет больше now
		for next.Before(now) || next.Equal(now) {
			next = next.AddDate(0, 0, days)
		}
		return next.Format(DateFormat), nil

	case "y":
		if len(parts) != 1 {
			return "", fmt.Errorf("неправильный формат правила y")
		}
		next := date
		// Всегда добавляем год
		next = next.AddDate(1, 0, 0)
		// Повторяем, пока не станет больше now
		for next.Before(now) || next.Equal(now) {
			next = next.AddDate(1, 0, 0)
		}
		// Корректировка для 29 февраля
		if next.Month() == 2 && next.Day() == 29 && !isLeapYear(next.Year()) {
			next = time.Date(next.Year(), 3, 1, 0, 0, 0, 0, time.UTC)
		}
		return next.Format(DateFormat), nil

	default:
		return "", fmt.Errorf("неподдерживаемый формат правила: %s", rule)
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
