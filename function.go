package scraper

import "net/http"

// Handle is public
func Handle(w http.ResponseWriter, r *http.Request) {
    msg := "Hello World"
    w.Write([]byte((msg)))
}