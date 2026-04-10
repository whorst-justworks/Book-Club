package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//UnbufferedChannel()
	BufferedChannel()
}

func UnbufferedChannel() {
	ch := make(chan int) // unbuffered channel

	var wg sync.WaitGroup
	wg.Add(2)

	// Sender Go Routine
	go func() {
		defer wg.Done()
		fmt.Println("Sender: Trying to send...")
		ch <- 42 // This will block until the receiver go routine awakes
		fmt.Println("Sender: Sent 42")
	}()

	// Receiver Go Routine
	go func() {
		defer wg.Done()
		fmt.Println("Receiver: Sleeping for a long time...")
		time.Sleep(10 * time.Second) // Not ready for 10 seconds

		val := <-ch
		fmt.Println("Received:", val)
	}()
	wg.Wait()
}

func BufferedChannel() {
	ch := make(chan int, 2) // buffered channel

	var wg sync.WaitGroup
	wg.Add(2)

	// Sender Go Routine
	go func() {
		defer wg.Done()
		fmt.Println("Sender: Trying to send...")
		ch <- 42 // This will not block at all
		fmt.Println("Sender: Sent 42")
	}()

	// Receiver Go Routine
	go func() {
		defer wg.Done()
		fmt.Println("Receiver: Sleeping for a long time...")
		time.Sleep(10 * time.Second) // Not ready for 10 seconds

		val := <-ch
		fmt.Println("Received:", val)
	}()
	wg.Wait()
}

// An unbuffered channel enables synchronization. We have the guarantee that two goroutines will be in a
// known state: one receiving and another sending a message.

// A buffered channel doesn’t provide any strong synchronization. A sender goroutine can send a message and then
// continue its execution if the channel isn’t full. The only guarantee is that a goroutine won’t receive a message
// before it is sent. But this is only a guarantee because of causality (you don’t drink your coffee before you prepare it).

// What size do we use for buffered channels?
/*

1) While using a worker pooling-like pattern, meaning spinning a fixed number of goroutines that need to send data to a
shared channel. In that case, we can tie the channel size to the number of goroutines created.

2) When using channels for rate-limiting problems. For example, if we need to enforce resource utilization by bounding
the number of requests, we should set up the channel size according to the limit.
*/
