package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func incompleteHandler(w http.ResponseWriter, r *http.Request) {
	// create a new App Engine context from the HTTP request.
	ctx := appengine.NewContext(r)

	p := &Person{Name: "gopher", AgeYears: 5}

	// create a new complete key of kind Person.
	key := datastore.NewIncompleteKey(ctx, "Person", nil)
	// put p in the datastore.
	key, err := datastore.Put(ctx, key, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "gopher stored with key %v", key)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	//name := mux.Vars(r)["key"]
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "Person", "gopher", 0, nil)

	var p Person
	err := datastore.Get(ctx, key, &p)
	if err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, p)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var p []Person

	q := datastore.NewQuery("Person").
		Filter("AgeYears <=", 10)

	// and finally execute the query retrieving all values into p.
	_, err := q.GetAll(ctx, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, p)
}

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/complete", completeHandler)
	r.HandleFunc("/incomplete", incompleteHandler)
	r.HandleFunc("/query", queryHandler)
	r.HandleFunc("/get/{key}", getHandler)
	http.Handle("/", r)
}
