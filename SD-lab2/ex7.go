// ex 7
package main

import (
	"fmt"
	"time"
)

type Point struct {
	X, Y int
}

func robot(id string, toOther chan interface{}, observerDone chan string) {
	
	for i := 1; i <= 3; i++ {
		toOther <- Point{X: i * 10, Y: i * 5}
		time.Sleep(100 * time.Millisecond)
	}
	toOther <- "STOP"

	for {
		msg := <-toOther
		if msg == "STOP" {
			break
		}
	}

	observerDone <- id + " a terminat sarcina"
}

func main() {
	ch1 := make(chan interface{}, 10)
	ch2 := make(chan interface{}, 10)
	observerDone := make(chan string)

	go robot("Robot 1", ch1, observerDone)
	go robot("Robot 2", ch2, observerDone)

	fmt.Println("Simularea sistemului distribuit a început...")
}