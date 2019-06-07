package entities

import (
	"time"
)

// Page struct
type Page struct {
	URL         string `datastore:"-" goon:"id" json:"url"`
	CreatedAt time.Time      `json:"createdAt"`
}
