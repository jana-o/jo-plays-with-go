package main

import "fmt"

func main() {

	c := make(chan int)
	//send
	go foo(c)

	//receive - not a go routine bc it would block
	bar(c)

	fmt.Println("bye")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
