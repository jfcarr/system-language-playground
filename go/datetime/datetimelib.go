package main

import (
	"time"
)

type DateTime struct {
	InputDateTime string
	Month         string
	Day           string
	Year          string
	Hour          string
	Minute        string
	Second        string
	FormattedTime string
	TimeZone      string
	CurrentError  error
}

func (dt *DateTime) ParseInput() {
	t1, e := time.Parse(
		time.RFC3339,
		dt.InputDateTime)

	dt.Month = t1.Format("01")
	dt.Day = t1.Format("02")
	dt.Hour = t1.Format("03")
	dt.Minute = t1.Format("04")
	dt.Second = t1.Format("05")
	dt.Year = t1.Format("2006")
	dt.TimeZone = t1.Format("-07")

	dt.FormattedTime = t1.Format("3:04 PM")

	dt.CurrentError = e
}
