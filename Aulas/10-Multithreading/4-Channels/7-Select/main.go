package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

func main() {

	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {

		for {
			atomic.AddInt64(&i, 1)
			c1 <- Message{i, "Hello from RabbitMq"}
		}

	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			c2 <- Message{i, "Hello from Kafka"}
		}
	}()

	for {

		select {
		case msg := <-c1:
			fmt.Printf("Receive from RabbitMq: ID: %d - %s\n", msg.id, msg.msg)

		case msg := <-c2:
			fmt.Printf("Receive from Kafka: ID: %d - %s\n", msg.id, msg.msg)

		case <-time.After(time.Second * 3):
			println("Timeout")

		}

	}
}
