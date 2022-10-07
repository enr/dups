package main

import (
	"path/filepath"
	"testing"
	"time"
)

type durationTestCases struct {
	duration time.Duration
	out      string
}

var testcases = []durationTestCases{
	{0, "no time"},
	{500 * time.Millisecond, "500 ms"},
	{1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Millisecond, "1 hour 2 minutes 3 seconds 4 ms"},
}

func TestHumanizeDuration(t *testing.T) {

	for _, data := range testcases {
		actual := humanizeDuration(data.duration)
		if actual != data.out {
			t.Errorf(`duration="%v" got="%s" expected="%s"`, data.duration, actual, data.out)
		}
	}
}

func abs(p string) string {
	abspath, err := filepath.Abs(p)
	if err != nil {
		return filepath.FromSlash(p)
	}
	return filepath.FromSlash(abspath)
}
