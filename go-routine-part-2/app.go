package main

import (
	"fmt"
	"time"
)

func main() {

	pipe := make(chan bool)
	go printer("Hello - 1", pipe)
	go printer("Hello - 2", pipe)
	go printer("Hello - 3", pipe)
	go sleepPrinter("Hello - 1 late", pipe)

	// Use the same channel for different go routines and use the range keyword to read from the channel until it is closed
	for donePipe := range pipe {
		fmt.Println("Done with ", donePipe)
	}

}

func printer(data string, pipe chan bool) {
	fmt.Println(data)
	pipe <- true
}

func sleepPrinter(data string, pipe chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(data)
	pipe <- true
	// This tells the main function that no more data will be sent to the channel and it can stop waiting for data from the channel
	close(pipe)
}
