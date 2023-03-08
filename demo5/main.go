package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	exit := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")

			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed, leaving")

		close(ch)
	}()

	go func() {
		//  Example using select
		// The select is more useful with multiple channels.
		for {
			select {
			case v, open := <-ch:
				if !open {
					close(exit)
					return
				}

				fmt.Println(time.Now(), "received", v)
			default:
				fmt.Println("Noting is happening")
			}
		}

		//We could use the range below and comment out the for loop above as we are only using 1 channel
		// XXX: In cases where only one channel is used
		// for v := range ch {
		// 	fmt.Println(time.Now(), "received", v)
		// }

		// close(exit)
	}()

	fmt.Println(time.Now(), "waiting for everything to complete")

	<-exit

	fmt.Println(time.Now(), "exiting")
}
