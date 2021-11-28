package cron

import (
	"errors"
	"strconv"
	"strings"
)

var (
	days    = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	months  = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	special = []string{"*", ",", "-", "/"}
	order   = []string{"minute", "hour", "dayofmonth", "month", "day"}
)

type ParsedCron struct {
	Minute     int
	Hour       int
	DayOfMonth int
	Month      int
	DayOfWeek  int
}

func (pc *ParsedCron) Next() {}

type CronParser struct{}

func (cp *CronParser) ParseCron(cron string) {}

func (cp *CronParser) ValidateCron(cron string) error {
	cronSplit := strings.Split(cron, " ")
	if len(cronSplit) != 5 {
		// Current implementation does not allow seconds or explicit years
		return errors.New("too many arguments")
	}

	for index, val := range cronSplit {
		if val == "*" {
			continue
		}

		intVal, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		if order[index] == "minute" {
			if !(intVal >= 0 && intVal <= 60) {
				return errors.New("int parsing failed")
			}
		}

		switch order[index] {
		case "minute":
			if !(intVal >= 0 && intVal <= 60) {
				return errors.New("minute parsing failed")
			}
		case "hour":
			if !(intVal >= 0 && intVal <= 24) {
				return errors.New("hour parsing failed")
			}
		case "dayofmonth":
			if !(intVal >= 0 && intVal <= 31) {
				return errors.New("dayofmonth parsing failed")
			}
		case "month":
			if !(intVal >= 1 && intVal <= 12) {
				return errors.New("month parsing failed")
			}
		case "dayofweek":
			if !(intVal >= 0 && intVal <= 6) {
				return errors.New("dayofweek parsing failed")
			}
		}
	}

	return nil
}

func ExistIn(element string, slice []string) (int, bool) {
	for i, v := range slice {
		if element == v {
			return i, true
		}
	}
	return 0, false
}
