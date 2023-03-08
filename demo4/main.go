package main

import (
	"fmt"
	"time"
)

func main() {
	//Buffered Channel will have a capacity
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
		}
		/*
			Race condition exist as there could be cases where this message is not completed,
		*/

		fmt.Println(time.Now(), "all completed")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "waiting for messages")

	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)

	fmt.Println(time.Now(), "exiting")
}
