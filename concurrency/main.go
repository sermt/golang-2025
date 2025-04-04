package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
}

func main() {
	dones := make([]chan bool, 4)
	for i := 0; i < 4; i++ {
		dones[i] = make(chan bool)
	}
	//done := make(chan bool)
	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])
	for i := 0; i < 4; i++ {
		<-dones[i]
	}
}
