package main

import (
	"fmt"
	"sync"
)

func main() {
	//raceConditionProblem()
	//better()
	best()
}

// go run -race Chapter9/69/69.go
func raceConditionProblem() {
	// Length 0
	s := make([]int, 0)

	go func() {
		s = append(s, 1)
		fmt.Println(s)
	}()

	go func() {
		s = append(s, 1)
		fmt.Println(s)
	}()
}

// go run -race Chapter9/69/69.go
func better() {
	s := make([]int, 0)
	var var1 []int
	var var2 []int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		var1 = append(var1, 1)
	}()

	go func() {
		defer wg.Done()
		var2 = append(var2, 2)
	}()

	wg.Wait() // Wait for goroutines to finish
	s = append(s, var1...)
	s = append(s, var2...)
	fmt.Println(s)
}

// go run -race Chapter9/69/69.go
func best() {
	s := make([]int, 0)
	ch := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	// Multiple goroutines can send values to the channel
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- 1
	}()

	go func() {
		defer wg.Done()
		ch <- 2
	}()

	// Single goroutine that handles all appends to s
	// This eliminates race conditions on s
	go func() {
		for val := range ch {
			s = append(s, val)
		}
		done <- true // Signal when finished processing
	}()

	// Wait for all senders to finish, then close the channel
	wg.Wait()
	close(ch)

	// Wait for the appender to finish processing all values
	<-done

	fmt.Println("Final slice:", s)
}
