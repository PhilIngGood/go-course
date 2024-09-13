package main

import (
	"fmt"
	"time"
)

func greet(phrase string, channel chan bool) {
	fmt.Println("Hello!", phrase)
	channel <- true
}

func slowGreet(phrase string, channel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	channel <- true
	close(channel)
}

// go key word launch as a goroutine
func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	for range done {
		//fmt.Println(doneChan)
	}
}
