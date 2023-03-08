package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	messages := make(chan string)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			mu.Lock()                                                // acquire the mutex
			defer mu.Unlock()                                        // release the mutex when done
			messages <- fmt.Sprintf("Message from goroutine %d", id) // send a message to the channel
			wg.Done()
		}(i)
	}

	go func() {
		for msg := range messages {
			fmt.Println(msg) // print the message
		}
	}()

	wg.Wait()
	close(messages)
}

/*

In this example, we create a channel of strings and a mutex using the sync package. We then create 5 goroutines that send messages to the channel, while acquiring and releasing the mutex to ensure that only one goroutine can access the shared resource at a time. We also create a separate goroutine that receives messages from the channel and prints them to the console.

Note that we use a WaitGroup to wait for all the goroutines to finish before closing the channel, and we use the defer keyword to ensure that the mutex is always released even if a panic occurs.


*/
