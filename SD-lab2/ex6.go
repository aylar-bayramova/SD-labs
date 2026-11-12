// ex 6
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(name string, ch chan string) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
	ch <- name
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go robot("Robot 1", ch1)
	go robot("Robot 2", ch2)

	select {
	case winner := <-ch1:
		fmt.Printf("Câștigătorul cursei este: %s\n", winner)
	case winner := <-ch2:
		fmt.Printf("Câștigătorul cursei este: %s\n", winner)
	}
}