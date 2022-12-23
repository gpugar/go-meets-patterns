package alert

import (
	"fmt"
	"go-command-pattern/bus"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...any) string {
	sprint := func(args ...any) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
func Process(event bus.Event) {
	alert := event.Data.(Alert)
	switch alert.Priority {
	case Med:
		fmt.Printf(Teal("Received alert: %+v\n"), event.Data)
	case High:
		fmt.Printf(Yellow("Received alert: %+v\n"), event.Data)
	case Critical:
		fmt.Printf(Red("Received alert: %+v\n"), event.Data)
	default:
		fmt.Printf(Green("Received alert: %+v\n"), event.Data)
	}
}
