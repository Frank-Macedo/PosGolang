package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	go task("A",&wg)
	go task("b",&wg)

	wg.Wait()

}

func task(xpto string, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		fmt.Printf("%d Task %s is running\n", i, xpto)
	}

	wg.Done()
}
