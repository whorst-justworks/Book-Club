package _41

import "strings"


largeString := string(make([]byte, 1_000_000))

getMessageBad(largeString)

func getMessageBad(message string) string {
	// assume message is 10KB of data
	// We only need first 10 characters as ID
	return message[:10]


	// Problem: The entire 10KB byte array stays in memory
	// the substring still references the original array!
}


largeString := string(make([]byte, 1_000_000))
getMessageGood(largeString)

func getMessageGood(message string) string {
	id := message[:10]

	// Force a copy using string builder or conversion
	return strings.Clone(id)  // as of Go 1.18+ (best option)

	// Or alternative methods:
	// return string([]byte(id))           // Creates new backing array
	// return strings.Builder.String(id)   // Also creates copy
}