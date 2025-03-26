package utils

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

func IndexOfString(slice []string, target string) int {
	for i, value := range slice {
		if value == target {
			return i // Return the index if found
		}
	}
	return -1 // Return -1 if the target is not found
}

func ConvertStringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		slog.Error("Failed to convert string to int", ":", err)
		return 0, err
	}
	return int(i), err
}

func ConvertStringToDate(dateStr string, toBuddhist bool) (time.Time, error) {
	if strings.Contains(dateStr, "/") {
		parts := strings.Split(dateStr, "/")
		if len(parts) == 3 {
			day := parts[0]
			month := parts[1]
			buddhistYear := parts[2]
			year, err := strconv.Atoi(buddhistYear)
			if err != nil {
				return time.Time{}, fmt.Errorf("invalid year: %s", buddhistYear)
			}
			year -= 543

			convertedDateStr := fmt.Sprintf("%d-%s-%s", year, month, day)
			parsedDate, err := time.Parse("2006-01-02", convertedDateStr)
			if err != nil {
				return time.Time{}, err
			}

			if toBuddhist {
				parsedDate = parsedDate.AddDate(543, 0, 0)
			}
			return parsedDate, nil
		}
	}

	layouts := []string{
		"2006-01-02",
		"01/02/2006",
		"02-01-2006",
		"2006-01-02T15:04:05Z07:00",
	}

	for _, layout := range layouts {
		date, err := time.Parse(layout, dateStr)
		if err == nil {
			if toBuddhist {
				return date.AddDate(543, 0, 0), nil
			}
			return date, nil
		}
	}
	return time.Time{}, fmt.Errorf("failed to convert string to date: %s", dateStr)
}
