package main

import (
    "fmt"
    "github.com/sonomy/scraper"
    "google.golang.org/appengine"
    "net/http"
)

func main() {
    httpPort := 8080
    fmt.Printf("> listening on %v\n", httpPort)
    http.HandleFunc("/scraper", scraper.Handle)
    appengine.Main()
}
