package input

import (
	"fmt"
	"go-command-pattern/bus"
	"go-command-pattern/command"
)

func Process(event bus.Event) {
	factory := command.GetFactoryInstance()
	reading, err := factory.CreateReading([]byte(event.Data.(string)))
	if err != nil {
		fmt.Printf("Error deserializing event `%s`: %s\n", event.Data, err)
		return
	}
	if err := reading.Execute(); err != nil {
		fmt.Printf("Error executing event: %s\n", err)
	}
}
