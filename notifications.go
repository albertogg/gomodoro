package main

import (
	"fmt"

	"github.com/deckarep/gosx-notifier"
)

func notify(message string) {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Gomodoro"
	note.Sound = gosxnotifier.Basso

	err := note.Push()
	if err != nil {
		fmt.Println("Problems with the notifications")
	}
}
