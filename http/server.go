package http

import (
	"errors"
	"fmt"
	"go-command-pattern/bus"
	"io"
	"net/http"
)

const keyServerAddr = "serverAddr"

func processRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("got / request. body:\n%s\n", body)
	go bus.PublishTo([]string{"input"}, string(body))
}

func InitServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", processRequest)

	err := http.ListenAndServe(":13333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
