package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Allow users to specify which directory to serve, defaulting to the users Public directory
var directory_to_serve = flag.String("dir", filepath.Join(os.Getenv("HOME"), "Public"), "the static directory to serve")

func logRequests(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do stuff with the request here
		log.Println("Request logged")
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	port_to_serve := ":8080"

	log.Println("FileServer started on port", port_to_serve, "...")
	log.Println("Serving directory: ", *directory_to_serve)
	handler := logRequests(http.FileServer(http.Dir(*directory_to_serve)))
	log.Fatal(http.ListenAndServe(port_to_serve, handler))
}
