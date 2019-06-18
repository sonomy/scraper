package scraper

import (
    "log"
    "net/http"
    "context"
    "os"
    "google.golang.org/appengine"
    "cloud.google.com/go/datastore"
    // "github.com/mjibson/goon"
    // "github.com/sonomy/scraper/entities"
)

type MyEntity struct {
	A int
	K *datastore.Key `datastore:"__key__"`
}

// Handle is public
func Handle(w http.ResponseWriter, r *http.Request) {
    projID := os.Getenv("GCP_PROJECT")
    if projID == "" {
        log.Fatal(`You need to set the environment variable "GCP_PROJECT"`)
    }

    var ctx context.Context
    host := os.Getenv("DATASTORE_EMULATOR_HOST")
	if false {
        log.Printf(`DATASTORE_EMULATOR_HOST alive: %v`, host)
        ctx = appengine.NewContext(r)
    } else {
        ctx = context.Background()
    }

    client, err := datastore.NewClient(ctx, projID)
    if err != nil {
        log.Fatalf("Could not create datastore client: %v, %v", err, client)
    }

    k := datastore.NameKey("Entity", "stringID2", nil)
    e := MyEntity{A: 12}
    if _, err := client.Put(ctx, k, &e); err != nil {
        log.Fatalf("Could not create entity: %v, %v", err, e)
    }

    var entities []MyEntity
    q := datastore.NewQuery("Entity").Limit(5)
    // q := datastore.NewQuery("Entity").Filter("A =", 12).Limit(1)
    if _, err := client.GetAll(ctx, q, &entities); err != nil {
        log.Fatalf("Could not get entity: %v, %v", err, e)
    }

    log.Println(entities)

    urls, ok := r.URL.Query()["url"]

    if !ok || len(urls[0]) < 1 {
        log.Println("Url Param 'url' is missing!!")
        return
    }

    url := urls[0]

    // obj := &entities.Page{
    //     URL: "http://test.com",
    // }

    // g := goon.FromContext(ctx)
    // _, err := g.Put(obj)
    // if err != nil {
    //     log.Fatalf("Put failed. %v", err)
    // }


    //_msg := "Hello World"
    w.Write([]byte((url)))
}