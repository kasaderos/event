package main

import (
	"fmt"

	"github.com/kasaderos/event"
)

func main() {
	// enable logging
	event.Logging = true

	manager := event.NewManager()

	event1 := event.Event{
		"Do job!",
		nil,
	}
	event2 := event.Event{
		"Pass event1 to event3",
		nil,
	}
	event3 := event.Event{
		"Job finished",
		func() {
			fmt.Println("Doing job")
		},
	}

	manager.Bind(event1, event2)
	manager.Bind(event2, event3)
	manager.Initiate(event1)
}
