//ex 3
package main

import (
	"fmt"
)

func sender(ch chan string) {
	ch <- "primul mesaj"
	ch <- "al doilea mesaj"
}

func main() {
	ch := make(chan string)
	go sender(ch)

	msg1 := <-ch
	msg2 := <-ch

	fmt.Println(msg1)
	fmt.Println(msg2)
}