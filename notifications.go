package main

import (
	"fmt"

	"github.com/deckarep/gosx-notifier"
)

// this method uses de gosx-notifier api to send a message to the
// osx notification center
func notify(message string) {
	note := gosxnotifier.NewNotification(message)
	note.Title = "Gomodoro"
	note.Sound = gosxnotifier.Basso

	err := note.Push()
	if err != nil {
		fmt.Println("Problems with the notifications")
	}
}
