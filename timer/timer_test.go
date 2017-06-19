package timer

import (
	"testing"
	"time"
)

func TestStartTimeIsBeforeNow(t *testing.T) {
	var testTime = Start()
	if !testTime.Before(time.Now()) {
		t.Error("Expected start time to be before now")
	}
}

func TestStopTimeIsAfterStart(t *testing.T) {
	var testTime = Stop()
	if !testTime.After(startTime) {
		t.Error("Expected stop time to be after start time")
	}
}
