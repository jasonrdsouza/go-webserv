package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func logRequests(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //do stuff with the request here
    log.Println("Request logged")
    h.ServeHTTP(w, r)
  })
}

func main() {
  path_to_serve := filepath.Join(os.Getenv("HOME"), "Public")
  port_to_serve := ":8080"

  log.Println("FileServer started on port", port_to_serve, "...")
  log.Println("Serving directory: ", path_to_serve)
  handler := logRequests(http.FileServer(http.Dir(path_to_serve)))
  log.Fatal(http.ListenAndServe(port_to_serve, handler))
}
