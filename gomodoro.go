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

var pomodori, shortBreak, largeBreak, pomodoriRun int

func init() {
	flag.IntVar(&pomodori, "p", 25, "Pomodoros work time (minutes)")
	flag.IntVar(&shortBreak, "s", 5, "Short break time (minutes)")
	flag.IntVar(&largeBreak, "l", 30, "Large break time (minutes)")
	flag.IntVar(&pomodoriRun, "r", 4, "Pomodori Runs, How many pomodoro runs until large break")
}

func sleepTimer(t int, message string) {
	notify(message)
	fmt.Println(message)
	time.Sleep(time.Duration(t) * time.Minute)
}

func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [options]\n\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Options:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = show_usage
	flag.Parse()

	for i := 1; i <= pomodoriRun; i++ {
		fmt.Println("Pomodoro run #", i)
		sleepTimer(pomodori, PomodoriMessage)

		if i == (pomodoriRun) {
			fmt.Print("Large break: ")
			sleepTimer(largeBreak, LargeMessage)
		} else {
			fmt.Print("Short break: ")
			sleepTimer(shortBreak, ShortMessage)
		}
	}

}
