// ex 1
package main

import (
	"fmt"
)
func worker(ch chan string){
	ch <- "Buna"
}
func main() {
	ch := make(chan string)
	go worker(ch)

	msg := <- ch
	fmt.Println(msg)
}