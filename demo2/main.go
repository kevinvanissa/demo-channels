package main

import "fmt"

// This program panics because there is no goroutine
// outside of `main` interacting with the `ch` channel:
//
// fatal error: all goroutines are asleep - deadlock!

// There are no go routines interacting with the chanel outside of main,
// therefore this will result in a deadlock
func main() {
	var ch chan int
	ch = make(chan int)
	// Two lines above can be also replaced with
	// ch := make(chan int)

	ch <- 10

	v := <-ch
	fmt.Println("received", v)
}
