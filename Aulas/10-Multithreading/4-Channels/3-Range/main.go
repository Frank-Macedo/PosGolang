package main

import "fmt"

func main() {

	ch := make(chan int)

	go publish(ch)

	reader(ch)
}

func reader(ch chan int) {

	for x := range ch {
		fmt.Printf("Reveived %d\n", x)
	}

}

func publish(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- 1
	}
	close(ch)

}
