// ex 5
package main

import (
	"fmt"
)

func send(ch chan string) {
	ch <- "Salut, lume!"
}

func receive(ch chan string) string {
	return <-ch
}

func main() {
	ch := make(chan string)
	go send(ch)

	result := receive(ch)
	fmt.Println(result)
}