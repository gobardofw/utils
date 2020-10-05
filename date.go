package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jalaali/go-jalaali"
	"github.com/mavihq/persian"
)

// StringToTime parse string date to date
func StringToTime(format string, date string) (time.Time, error) {
	t, err := time.Parse(format, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// JalaliToTime convert jalali date string to time from 2006-01-02 format
func JalaliToTime(jDate string) (time.Time, error) {
	rx := regexp.MustCompile(`^\d{4}(-|\/)(\d{1,2})(-|\/)(\d{1,2})$`)
	if !rx.MatchString(jDate) {
		return time.Time{}, errors.New("Inalid date format")
	}

	jDate = strings.ReplaceAll(jDate, "/", "-")
	parts := strings.Split(jDate, "-")

	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	if !jalaali.IsValidDate(year, month, day) {
		return time.Time{}, errors.New("Inalid date")
	}

	y, m, d, err := jalaali.ToGregorian(year, jalaali.Month(month), day)
	if err != nil {
		return time.Time{}, err
	}
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	return t, nil
}

// TimeToJalali convert time to jalali date
func TimeToJalali(t time.Time, format string) (string, error) {
	jDate := jalaali.From(t)
	date, err := jDate.JFormat(format)
	if err != nil {
		return "", err
	}
	return persian.ToEnglishDigits(date), nil
}
