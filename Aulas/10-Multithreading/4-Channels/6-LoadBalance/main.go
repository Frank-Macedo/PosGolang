package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	qtdWorker := 10000

	for i := 0; i < qtdWorker; i++ {
		go worker(i, data)

	}

	for i := 0; i < 100000; i++ {
		data <- i
	}
}
