package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("todo")
}

func ping() {
	// TODO - time the request
	res, err := http.Head("https://golang.org")
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
