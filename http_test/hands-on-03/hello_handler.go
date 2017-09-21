package main

import (
	"fmt"
	"log"
	"net/http"
)

func paramHandler(w http.ResponseWriter, r *http.Request) {

	name := ""

	for _, values := range r.Form { // range over map
		r.Form.
			fmt.Println(values)
		//		if key == "name" {
		for _, value := range values { // range over []string
			if name == "" {
				name = value
			} else {
				name += "and " + value
			}
		}
		//		}
	}

	if name == "" {
		name = "friend"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/hello", paramHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
