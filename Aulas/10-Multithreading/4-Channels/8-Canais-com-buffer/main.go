package main

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	print(<-ch)
	print(<-ch)

}