//ex 4
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Procesarea este în curs...")
	time.Sleep(time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done 
	fmt.Println("Procesare terminată, funcția main se încheie.")
}