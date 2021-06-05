package main

import (
	"net/http"

	"github.com/burhanuday/basic-web-application/pkg/handlers"
)

const port = ":5000"

// main is the entrypoint
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(port, nil)
}
