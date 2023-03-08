package main

import (
	"fmt"
	"runtime"
)

// Get the number of goroutines.
func main() {
	go hello()
	goRtns := runtime.NumGoroutine()
	fmt.Println("goroutines:", goRtns)
}

// go routing run concurrently so most time you will not see a result from hello() as main executes without waiting
//main is considered to be a go routine

// cool
func hello() {
	fmt.Println("it's most likely you will never see this")
}
