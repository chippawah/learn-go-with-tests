package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})
		want := "3\n2\n1\nGo!"
		got := buffer.String()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before each print of count", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{write, sleep, write, sleep, write, sleep, write}
		calls := spySleepPrinter.Calls
		if !reflect.DeepEqual(want, calls) {
			t.Errorf("wanted calls %v got calls %v", want, calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleeperTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleeperTime, sleep: spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleeperTime {
		t.Errorf("should have slept for %v but slept for %v", sleeperTime, spyTime.durationSlept)
	}
}
