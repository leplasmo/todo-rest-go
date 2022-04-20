package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// define the command-line parameters
	host := flag.String("h", "localhost", "Server host")
	port := flag.Int("p", 8080, "Server port")
	todoFile := flag.String("f", "todoServer.json", "todo JSON file")
	flag.Parse()

	// create an http server instance
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      newMux(*todoFile),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// run the server
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
