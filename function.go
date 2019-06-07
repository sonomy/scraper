package scraper

import (
    "log"
    "net/http"
    // "context"
    // "os"
    // "google.golang.org/appengine/datastore"
)

// Handle is public
func Handle(w http.ResponseWriter, r *http.Request) {

    // projID := os.Getenv("DATASTORE_PROJECT_ID")
	// if projID == "" {
	// 	log.Fatal(`You need to set the environment variable "DATASTORE_PROJECT_ID"`)
	// }
	// // [START datastore_build_service]
	// ctx := context.Background()
	// client, err := datastore.NewClient(ctx, projID)
	// // [END datastore_build_service]
	// if err != nil {
	// 	log.Fatalf("Could not create datastore client: %v", err)
	// }


    urls, ok := r.URL.Query()["url"]

    if !ok || len(urls[0]) < 1 {
        log.Println("Url Param 'url' is missing")
        return
    }

    url := urls[0]
    //_msg := "Hello World"
    w.Write([]byte((url)))
}