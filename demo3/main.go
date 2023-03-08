package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	//This an anonymous function
	go func() {
		fmt.Println(time.Now(), "taking a nap")

		//sleep to see what is happening
		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	fmt.Println(time.Now(), "waiting for message")
	//Channels by default sends and receives block until the other side is ready
	v := <-ch

	fmt.Println(time.Now(), "received", v)
}
