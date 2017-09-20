package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	// try changing the value of this url
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 404 {
		fmt.Println("Couldn't find that page m8")
	} else {
		if b, err := ioutil.ReadAll(res.Body); err == nil {
			fmt.Println(string(b))
		}
		res.Body.Close()
	}
}
