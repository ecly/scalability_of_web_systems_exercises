// code from: https://github.com/jf87/scalable_web_systems/blob/master/session-02/hands-on-04/README.md
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	AgeYears int    `json:"age_years"`
}

func encode() {
	p := Person{"gopher", 5}

	// create an encoder that will write on the standard output.
	enc := json.NewEncoder(os.Stdout)
	// use the encoder to encode p, which could fail.
	err := enc.Encode(p)
	// if it failed, log the error and stop execution.
	if err != nil {
		log.Fatal(err)
	}
}

func decode() {
	// create an empty Person value.
	var p Person

	// create a decoder reading from the standard input.
	dec := json.NewDecoder(os.Stdin)
	// use the decoder to decode a value into p.
	err := dec.Decode(&p)
	// if it failed, log the error and stop execution.
	if err != nil {
		log.Fatal(err)
	}
	// otherwise log what we decoded.
	fmt.Printf("decoded: %#v\n", p)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{"gopher", 5}

	// set the Content-Type header.
	w.Header().Set("Content-Type", "application/json")

	// encode p to the output.
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	if err != nil {
		// if encoding fails, create an error page with code 500.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var p Person

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name is %v and age is %v", p.Name, p.AgeYears)
}

func main() {
	http.HandleFunc("/", decodeHandler)
	http.ListenAndServe(":8080", nil)
}
