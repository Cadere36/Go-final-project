package scheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat is empty")
	}

	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}
	parts := strings.Split(repeat, " ")
	switch parts[0] {
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid repeat")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}
		if days > 400 || days <= 0 {
			return "", fmt.Errorf("invalid days")
		}
		for {
			date = date.AddDate(0, 0, days)
			if afterNow(date, now) {
				break
			}
		}
	case "w":
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid weeks")
		}
		week, err := parseWeekdays(parts[1])
		if err != nil {
			return "", err
		}
		if afterNow(now, date) {
			date = now
		}
		for {
			date = date.AddDate(0, 0, 1)
			wd := int(date.Weekday())
			if wd == 0 {
				wd = 7
			}
			if afterNow(date, now) && week[wd] {
				break
			}
		}
	case "m":
		if len(parts) != 2 && len(parts) != 3 {
			return "", fmt.Errorf("invalid months")
		}
		days, last, preLast, err := parseMonthDays(parts[1])
		if err != nil {
			return "", err
		}
		var months [13]bool
		if len(parts) == 3 {
			months, err = parseMonths(parts[2])
			if err != nil {
				return "", err
			}
		} else {
			for i := 1; i <= 12; i++ {
				months[i] = true
			}
		}
		if afterNow(now, date) {
			date = now
		}
		for i := 0; i < 3650; i++ {
			date = date.AddDate(0, 0, 1)
			dayOK := days[date.Day()] || (last && isLastDay(date)) || (preLast && isPreLastDay(date))
			if afterNow(date, now) && months[int(date.Month())] && dayOK {
				return date.Format(DateFormat), nil
			}
		}
		return "", fmt.Errorf("no matching date")

	default:
		return "", fmt.Errorf("invalid format")
	}
	return date.Format(DateFormat), nil

}

func NormalizeDate(date, repeat string, now time.Time) (string, error) {

	if date == "" {
		date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, date)
	if err != nil {
		return "", err
	}
	var next string
	if repeat != "" {
		next, err = NextDate(now, date, repeat)
		if err != nil {
			return "", err
		}
	}
	if afterNow(now, t) {
		if repeat == "" {
			return now.Format(DateFormat), nil
		}
		return next, nil

	}
	return date, nil
}

func afterNow(date, now time.Time) bool {
	return date.Format(DateFormat) > now.Format(DateFormat)
}

func isLastDay(d time.Time) bool {
	return d.AddDate(0, 0, 1).Month() != d.Month()
}

func isPreLastDay(d time.Time) bool {
	return d.AddDate(0, 0, 1).Month() == d.Month() && d.AddDate(0, 0, 2).Month() != d.Month()
}

func parseWeekdays(s string) ([8]bool, error) {
	var week [8]bool

	parts := strings.Split(s, ",")

	for _, v := range parts {
		n, err := strconv.Atoi(v)
		if err != nil {
			return week, err
		}
		if n < 1 || n > 7 {
			return week, fmt.Errorf("invalid value days in week")
		}
		week[n] = true
	}

	return week, nil
}

func parseMonths(s string) ([13]bool, error) {
	var months [13]bool

	parts := strings.Split(s, ",")

	for _, v := range parts {
		n, err := strconv.Atoi(v)
		if err != nil {
			return months, err
		}
		if n < 1 || n > 12 {
			return months, fmt.Errorf("invalid month number")
		}
		months[n] = true
	}

	return months, nil
}

func parseMonthDays(s string) ([32]bool, bool, bool, error) {
	var days [32]bool
	var last, preLast bool

	parts := strings.Split(s, ",")

	for _, v := range parts {
		n, err := strconv.Atoi(v)
		if err != nil {
			return days, false, false, err
		}
		switch {
		case n == -1:
			last = true
		case n == -2:
			preLast = true
		case n >= 1 && n <= 31:
			days[n] = true
		default:
			return days, false, false, fmt.Errorf("invalid day of month: %d", n)
		}
	}
	return days, last, preLast, nil
}
