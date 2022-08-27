package main

import (
	"fmt"
	"math/rand"
	"time"
)

// producerCount >= 1, consumerCount >= 1
const producerCount int = 2
const consumerCount int = 5

// bufferSize >= 0 (>= 1 for Dijkstra's assumption)
var bufferSize int = 2
var buffer = make(chan int, bufferSize)
var runforever = make(chan bool)

// sleep is for the output to be nicer, it is not a vital part of the algorithm
func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
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
		sleep()
		num := <-buffer
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
	// stop execution with keyboard interrupt: ctrl + c
	<-runforever
}
