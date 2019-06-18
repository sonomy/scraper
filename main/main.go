package main

import (
    "fmt"
    "github.com/sonomy/scraper"
    "google.golang.org/appengine"
    "net/http"
)

// @see https://gist.github.com/hoitomt/c0663af8c9443f2a8294

func main() {
    httpPort := 8080
    fmt.Printf("> listening on %v\n", httpPort)
    http.HandleFunc("/scraper", scraper.Handle)
    appengine.Main()
}
