package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func doGet() {
	req, err := http.NewRequest("GET", "https://http-methods.appspot.com/Hungary/", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}

	q := req.URL.Query()
	q.Add("v", "true")
	req.URL.RawQuery = q.Encode()

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed:\n %v\n", err)
	} else if b, err := ioutil.ReadAll(res.Body); err == nil {
		fmt.Println(string(b))
	}
	res.Body.Close()
}
