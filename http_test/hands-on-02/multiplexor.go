package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintln(w, "Hello", name)
}

func startMultiplexor() {
	r := mux.NewRouter()
	// match only GET requests on /product/
	r.HandleFunc("/hello/{name}", sayHello).Methods("GET")

	// handle all requests with the Gorilla router.
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
