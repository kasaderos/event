## Event Manager

### Installation

```
go get github.com/kasaderos/event
```

### Example:

```
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
		"Event passed",
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
```

### Output:

```
Doing job
2023-04-09T10:34:14+06:00 [events] Do job! -> Event passed -> *Job finished
```
