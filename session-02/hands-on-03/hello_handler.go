package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type textHandler struct {
	h func(http.ResponseWriter, *http.Request) (error, int)
}

func (t textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err, code := t.h(w, r)
	w.WriteHeader(code)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, strconv.Itoa(code))
	// Set the content type
	w.Header().Set("Content-Type", "text/plain")
}

func bodyHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	b, err := ioutil.ReadAll(r.Body)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
		err = errors.New("Terrible request")
	}
	body := string(b)
	if body == "" {
		code = http.StatusBadRequest
		err = errors.New("Empty body")
	}
	return err, code
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	name := ""

	r.ParseForm()
	names := r.Form["name"]
	// if multiple name values, say hello to them all
	for _, value := range names {
		if name == "" {
			name = value
		} else {
			name += " and " + value
		}
	}

	if name == "" {
		name = "friend"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.Handle("/hello", textHandler{bodyHandler})
	http.ListenAndServe(":8080", nil)
}
