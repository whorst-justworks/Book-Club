package main

import (
	"fmt"
	"sync"
	"time"
)

// Donation checker function using mutex
func listing1() {
	type Donation struct {
		mu      sync.RWMutex
		balance int
	}
	donation := &Donation{}

	// Donation check goroutine
	// We keep looping until we meet our donation goal
	// This implementation requires many CPU cycles having to constantly lock and
	// unlocks to check if the donation goal is met
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
}

// Now lets try using a channel
func listing2() {
	type Donation struct {
		balance int
		ch      chan int
	}

	donation := &Donation{ch: make(chan int)}

	// Listener goroutines
	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}
	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}
}

// Potential output here:
// $11 goal reached
// $15 goal reached
// Why?

// Using sync.Cond (Condition Variable)
func listing3() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	// sync.Cond relies on a mutex and will coordinate locks and unlocks on broadcast
	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			// Wait waits for a conditional broadcast will continue execution
			// The steps on Wait():
			// 1. Unlock the mutex (donation.cond.L.Unock())
			// 2. Suspend the goroutine and wait for notification(donation.cond.Brodcast)
			// 3. Lock the mutex for operation (donation.cond.L.Lock())
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		// Broadcast notifies all active goroutines waiting
		// There's also Signal which behaves similarly to chan struct() and will only notify a single goroutine
		// Final Note: sync.Cond does not buffer like a channel so if no goroutines are listening, the message gets lost.
		donation.cond.Broadcast()
	}
}
