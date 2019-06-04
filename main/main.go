package main

import (
    "github.com/sonomy/scraper"
    "net/http"
)

func main() {
    http.HandleFunc("/scrape", scraper.Handle)
    http.ListenAndServe(":8082", nil)
}