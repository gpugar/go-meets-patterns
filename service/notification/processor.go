package notification

import (
	"fmt"
	"go-command-pattern/bus"
)

func Process(event bus.Event) {
	fmt.Printf("Received notification [topic: %s]: %+v\n", event.Topic, event.Data)
}
