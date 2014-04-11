package main

import (
    "log"
    //"net/http"
    "os"
    "path/filepath"
)


func choose_logs(path string, info os.FileInfo, err error) error {
  if err != nil {
    log.Println("Error: ", err)
  } else if info.Name() == "summary.log" {
    log.Println("Found: ", info.Name())
  } else {
    log.Println("Skipped: ", info.Name())
  }
  return nil
}

func main() {
  base_path := filepath.Join(os.Getenv("HOME"), "Desktop", "test")

  log.Println("Serving directory: ", base_path)
  err := filepath.Walk(base_path, choose_logs)
  if err != nil {
    log.Fatalln(err)
  }
}
