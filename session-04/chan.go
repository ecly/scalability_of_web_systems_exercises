package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hola!"

	}()
	s := <-ch
	fmt.Println(s)
}
