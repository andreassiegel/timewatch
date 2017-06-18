package timer

import (
	"time"
	"fmt"
)

var startTime time.Time
var stopTime time.Time

type Workday struct {
	date     time.Time
	weekday  time.Weekday
	start    time.Time
	end      time.Time
	duration time.Duration
}

func Start() time.Time {
	startTime = time.Now()
	fmt.Printf("Start time: %v \n", startTime)
	return startTime
}

func Stop() time.Time {
	stopTime = time.Now()
	fmt.Printf("End time: %v \n", stopTime)
	return stopTime
}

func Summary() Workday {

	return Workday{
		date: timeToDate(startTime),
		weekday: startTime.Weekday(),
		start: startTime,
		end: stopTime,
		duration: stopTime.Sub(startTime)}
}

func timeToDate(timestamp time.Time) time.Time {

	year, month, day := timestamp.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, timestamp.Local().Location())
}

func timeToDateString(timestamp time.Time) string {

	return timestamp.Format("2006-01-02")
}

func timeToTimeString(timestamp time.Time) string {

	return timestamp.Format("15:04:05")
}

func (w Workday) String() string {

	return fmt.Sprintf("%v;%v;%v;%v;%v;", timeToDateString(w.date), w.weekday.String(), timeToTimeString(w.start), timeToTimeString(w.end), w.duration)
}
