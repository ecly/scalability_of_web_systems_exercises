package app

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Person struct {
	Name     string `json:"name"`
	AgeYears int    `json:"ageyears"`
}

func completeHandler(w http.ResponseWriter, r *http.Request) {
	// create a new App Engine context from the HTTP request.
	ctx := appengine.NewContext(r)

	p := &Person{Name: "gopher", AgeYears: 5}

	// create a new complete key of kind Person and value gopher.
	key := datastore.NewKey(ctx, "Person", "gopher", 0, nil)
	// put p in the datastore.
	key, err := datastore.Put(ctx, key, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "gopher stored with key %v", key)
}

func init() {
	http.HandleFunc("/hello", completeHandler)
}
