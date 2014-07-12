package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mitchellh/colorstring"
)

const (
	PomodoriMessage = "Starting Pomodoro Run"
	ShortMessage    = "Starting Short Break"
	LargeMessage    = "Starting Large Break"
)

var (
	pomodori, shortBreak, largeBreak, pomodoriRun int
	start                                         time.Time = time.Now()
)

// initialize the flags/options for the command line
func init() {
	flag.IntVar(&pomodori, "p", 25, "Pomodoros work time (minutes)")
	flag.IntVar(&shortBreak, "s", 5, "Short break time (minutes)")
	flag.IntVar(&largeBreak, "l", 30, "Large break time (minutes)")
	flag.IntVar(&pomodoriRun, "r", 4, "Pomodori Runs, How many pomodoro runs until large break")
}

func sleepTimer(t int, message string) {
	// notify is a function that lives in notification.go
	notify(message)
	fmt.Println(message)
	time.Sleep(time.Duration(t) * time.Minute)
}

// pretty prints the usage of the gomodoro command when a bad flag is used
func showUsage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [options]\n\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Options:\n")
	flag.PrintDefaults()
}

func elapsedTime(start time.Time) time.Duration {
	elapsed := time.Since(start)

	return elapsed
}

func catchUserInterruption(start time.Time) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf(colorstring.Color("Awww, your total pomodoro time was: %v, [red]%v\n"), elapsedTime(start), sig)
			os.Exit(2)
		}
	}()
}

func main() {
	flag.Usage = showUsage
	flag.Parse()

	catchUserInterruption(start)

	fmt.Println(colorstring.Color("[green]Start time:"), start.Format(time.RFC3339))
	for i := 1; i <= pomodoriRun; i++ {
		fmt.Printf("Run #%v\n", i)
		sleepTimer(pomodori, PomodoriMessage)

		if i%4 == 0 || i == (pomodoriRun) {
			sleepTimer(largeBreak, LargeMessage)
		} else {
			sleepTimer(shortBreak, ShortMessage)
		}
	}

	fmt.Println("Well done, your total pomodoro time was:", elapsedTime(start))
}
