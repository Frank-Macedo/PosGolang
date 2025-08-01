package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	reader(ch, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {

	for x := range ch {
		fmt.Printf("Reveived %d\n", x)
		wg.Done()
	}

}

func publish(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- 1
	}
	close(ch)

}
