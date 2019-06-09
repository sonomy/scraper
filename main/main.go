package main

import (
    "log"
    "fmt"
    "github.com/sonomy/scraper"
    "net/http"
)

// @see https://gist.github.com/hoitomt/c0663af8c9443f2a8294

func main() {
    httpPort := 8082
    fmt.Printf("> listening on %v\n", httpPort)
    http.HandleFunc("/scraper", scraper.Handle)
    http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}