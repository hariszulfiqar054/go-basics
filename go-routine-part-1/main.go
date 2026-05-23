package main

import (
	"fmt"
	"time"
)

func main() {

	// First parameter is the channel creation and 2nd one is the type of data it will transfer
	data := make(chan bool)

	greet("Hello - 1")
	greet("Hello - 2")
	greet("Hello - 3")
	go slowGreet("Hello - 1 late", data)
	<-data
}

func greet(value string) {
	fmt.Println(value)
}

func slowGreet(value string, data chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(value)
	// Putting true in the channel to indicate that the slowGreet function has completed its execution
	data <- true
}
