package scraper

import (
    "log"
    "net/http"
)

// Handle is public
func Handle(w http.ResponseWriter, r *http.Request) {
    urls, ok := r.URL.Query()["url"]

    if !ok || len(urls[0]) < 1 {
        log.Println("Url Param 'url' is missing")
        return
    }

    url := urls[0]
    //_msg := "Hello World"
    w.Write([]byte((url)))
}