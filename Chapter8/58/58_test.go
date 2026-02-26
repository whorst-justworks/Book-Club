package _58

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func listing1() int {
	// go test -race -run Test_Listing1 ./Chapter8/58 2>&1 | head -40
	i := 0
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		i++

	}()

	go func() {
		defer wg.Done()
		i++
	}()
	wg.Wait()
	return i
}

func Test_Listing1(t *testing.T) {
	for range 150 {
		fmt.Println(listing1())
	}

}

func listing2() int64 {
	var i int64
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	wg.Wait()
	return i
}

func Test_Listing2(t *testing.T) {
	//An atomic check-and-set operation is a synchronization primitive that reads a memory location, compares its value to an expected value, and if they match, updates it
	//to a new value - all as a single indivisible operation that cannot be interrupted by other threads. This ensures thread-safe updates in concurrent programs by
	//preventing race conditions where multiple threads might try to modify the same value simultaneously.
	for range 5 {
		result := listing2()
		fmt.Println(result)
	}
}

func listing3() int {
	//Is this data race free?
	i := 0
	mutex := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func(i *int) {
		defer wg.Done()

		mutex.Lock()
		*i++
		mutex.Unlock()
	}(&i)

	go func(i *int) {
		defer wg.Done()

		mutex.Lock()
		*i++
		mutex.Unlock()
	}(&i)

	wg.Wait()
	return i
}

func Test_Listing3(t *testing.T) {
	for range 5 {
		result := listing3()
		fmt.Println(result)
	}
}

func listing4() int {
	//Is this data race free?
	i := 0
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch
	return i
}

func Test_Listing4(t *testing.T) {
	for range 5 {
		result := listing4()
		fmt.Println(result)
	}
}

func listing5() int {
	//Is this data race free?
	var i int
	mutex := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		i = 1
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		i = 2
	}()

	wg.Wait()
	return i
}
func Test_Listing5(t *testing.T) {
	// Not deterministic, i will be either 1 or 2
	for range 5 {
		result := listing5()
		fmt.Println(result)
	}
}
