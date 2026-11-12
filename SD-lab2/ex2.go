// ex 2
package main

import (
	"fmt"
)

func worker(ch chan string) {
	msg := <-ch
	fmt.Println("Mesaj primit:", msg)
}

func main() {
	ch := make(chan string)
	go worker(ch)

	ch <- "salut"
}