package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// producerCount >= 1, consumerCount >= 1
	const producerCount int = 2
	const consumerCount int = 5

	// bufferSize >= 0 (>= 1 for Dijkstra's assumption)
	var bufferSize int = 2
	var buffer = make(chan int, bufferSize)
	var runForever = make(chan bool)

	// launch producer goroutines
	for i := 0; i < producerCount; i++ {
		go func(id int) {
			for {
				num := rand.Intn(100)
				fmt.Printf("Producer[%d] -> %d\n", id, num)
				buffer <- num
			}
		}(i + 1)
	}

	// launch consumer goroutines
	for i := 0; i < consumerCount; i++ {
		go func(id int) {
			for {
				num := <-buffer
				fmt.Printf("\t\t\t%2d -> Consumer[%d]\n", num, id)
			}
		}(i + 1)
	}
	// stop execution with keyboard interrupt: ctrl + c
	<-runForever
}
