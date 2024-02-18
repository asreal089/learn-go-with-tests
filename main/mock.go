package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "GO!"
	countdownStart = 3
)

func main() {
	sleeper := &DefaulfSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()
	}
	fmt.Println(finalWord)
}

type DefaulfSleeper struct{}

func (d *DefaulfSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
