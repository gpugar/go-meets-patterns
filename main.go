package main

import (
	"go-command-pattern/bus"
	"go-command-pattern/http"
	"go-command-pattern/service/alert"
	"go-command-pattern/service/input"
	"go-command-pattern/service/notification"
)

var inputChannel = make(chan bus.Event)
var simpleStdoutChannel = make(chan bus.Event)
var shinyStdoutChannel = make(chan bus.Event)

func subscribe() {
	busInstance := bus.GetInstance()
	// we can have topics that send to multiple channels
	busInstance.Subscribe("input", inputChannel)
	busInstance.Subscribe("input", simpleStdoutChannel)

	busInstance.Subscribe("notify", simpleStdoutChannel)
	busInstance.Subscribe("alert", shinyStdoutChannel)
}

func runRouter() {
	for {
		select {
		case d := <-inputChannel:
			go input.Process(d)
		case d := <-simpleStdoutChannel:
			go notification.Process(d)
		case d := <-shinyStdoutChannel:
			go alert.Process(d)
		}
	}
}

func main() {
	go http.InitServer()
	subscribe()
	runRouter()
}
