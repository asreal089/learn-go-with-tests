package main

import (
	"bytes"
	"testing"
	"time"
)

func testCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.duration != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.duration)
	}
}
