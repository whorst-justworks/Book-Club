package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Time time.Time
}

func badComparison() {
	t := time.Now()
	var event1 = Event{
		Time: t,
	}
	b, err := json.Marshal(event1)
	if err != nil {

	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
	}
	fmt.Println("is Time Equal ", event1 == event2)
	fmt.Println("event1 ", event1)
	fmt.Println("event2 ", event2)
}

// Time in Go contains references to both the Monotonic time and Wall time, which is confusing
// When we unmarshal json, the time.Time field takes away the monotonic time portion of the time field
//and leaves the wall time, seen in event2

// The m=+0.000070042 is the monotonic clock reading, measured in seconds since the process started. The
//+ simply means it is a positive offset from time zero (process start) — it has been 0.000070042 seconds
// (~70 microseconds) since your program began running when that timestamp was captured.

func goodComparison() {
	fmt.Println()
	t := time.Now()
	var event1 = Event{
		Time: t,
	}
	b, err := json.Marshal(event1)
	if err != nil {

	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
	}
	fmt.Println("is Time Equal ", event1.Time.Equal(event2.Time))
	fmt.Println("event1 ", event1)
	fmt.Println("event2 ", event2)
}

func main() {
	badComparison()
	goodComparison()

}
