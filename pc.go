// Producer-consumer, message passing solution
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// bufferSize >= 1, to fulfill Dijkstra's problem description
var bufferSize int = 1
var buffer = make(chan int, bufferSize)
var runforever = make(chan bool)

// producers and consumers count
const producerCount int = 2
const consumerCount int = 5

// simulating work and making output observable
func sleep() {
	time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
}

func producer(id int) {
	for {
		sleep()
		num := rand.Intn(100)
		buffer <- num
	}
}

func consumer(id int) {
	for {
		num := <-buffer
		sleep()
		fmt.Printf("Consumer%d: %d\n", id, num)
	}
}

func main() {
	for i := 0; i < producerCount; i++ {
		go producer(i + 1)
	}

	for i := 0; i < consumerCount; i++ {
		go consumer(i + 1)
	}
	// stop execution with keyboard interrupt (ctrl + c)
	<-runforever
}
