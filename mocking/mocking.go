package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countdownStart = 3
const finalWord = "Go!"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	fmt.Printf("Will sleep for %v", (cs.duration))
	time.Sleep(cs.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
