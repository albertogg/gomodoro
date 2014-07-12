package main

import (
	"flag"
	"fmt"
	"os"
	"time"
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

func main() {
	flag.Usage = showUsage
	flag.Parse()

	fmt.Println("Starting time is:", start.Format(time.RFC3339))

	for i := 1; i <= pomodoriRun; i++ {
		fmt.Println("Run #", i)
		sleepTimer(pomodori, PomodoriMessage)

		if i%4 == 0 || i == (pomodoriRun) {
			sleepTimer(largeBreak, LargeMessage)
		} else {
			sleepTimer(shortBreak, ShortMessage)
		}
	}

	fmt.Println("Well done, your total pomodoro time was:", elapsedTime(start))
}
