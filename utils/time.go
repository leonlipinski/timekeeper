package utils

import (
	"math"
	"fmt"
	"time"
)

// RoundToNearest15 rounds the given minutes to the nearest 15 minutes.
func RoundToNearest15(minutes int) int {
	if minutes == 0 {
		return 0
	}
	return int(math.Round(float64(minutes)/15.0)) * 15
}

// FormatDurationAsTime formats the given minutes into a string representation of hours and minutes.
// For example, 90 minutes will be formatted as "01:30".
// It returns a string in the format "HH:MM".
func FormatDurationAsTime(minutes int) string {
    hours := minutes / 60
    mins := minutes % 60
    return fmt.Sprintf("%02d:%02d", hours, mins)
}

// FormatAsMinutes converts hours and minutes into total minutes.
// For example, 1 hour and 30 minutes will be converted to 90 minutes.
// It returns the total minutes as an integer.
func FormatAsMinutes(hours int, minutes int) int {
	mins := minutes + (hours * 60)
	return mins
}

// FindMinMaxDates takes a slice of date strings and returns the earliest and latest date.
func FindMinMaxDates(dates []string) (string, string, error) {
	if len(dates) == 0 {
		return "", "", fmt.Errorf("date list is empty")
	}

	const layout = "2006-01-02"
	var minDate, maxDate time.Time

	for i, d := range dates {
		parsed, err := time.Parse(layout, d)
		if err != nil {
			return "", "", fmt.Errorf("invalid date format: %s", d)
		}

		if i == 0 || parsed.Before(minDate) {
			minDate = parsed
		}
		if i == 0 || parsed.After(maxDate) {
			maxDate = parsed
		}
	}

	return minDate.Format(layout), maxDate.Format(layout), nil
}
