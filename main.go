package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// user standard http library instead of external libraries

	// function to register route
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// send the response back to the client
		// Fprintf takes the "hello world" string, write it to w
		fmt.Fprint(w, "hello world")
	})

	// function to start server
	// nil because relying on the default multiplexer instead of creating one our own
	err := http.ListenAndServe("localhost:9000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
