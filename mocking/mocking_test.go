package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
	t.Run("verify calls", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("Not enough calls, got %q want %q", spySleeper.Calls, 4)
		}
	})

	t.Run("verify call order", func(t *testing.T) {
		countdown := &CountdownOperationSpy{}
		Countdown(countdown, countdown)
		want := []string{
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
		}

		if !reflect.DeepEqual(want, countdown.Calls) {
			t.Errorf("wanted calls %v got %v", want, countdown.Calls)
		}
	})
}
