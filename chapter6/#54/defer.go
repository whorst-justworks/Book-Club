package _54

import (
	"fmt"
	"os"
)

// Section #54: Not Handling Defer Errors

func writeToFileBad(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Problem: If Close() fails (e.g., disk full, can't flush buffer),
// we return nil suggesting success, but data may not be written!

// GOOD: Handling defer errors with named return value
func writeToFileGood(filename string, data []byte) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil && err == nil {
			// If no previous error, use Close error
			err = closeErr
		}
		// If err was already set, keep original error
	}()

	_, err = file.Write(data)
	return err
}

// GOOD Alternative: Wrap both errors if both fail
func writeToFileGoodWithWrapping(filename string, data []byte) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			if err != nil {
				// Both Write and Close failed - wrap both
				err = fmt.Errorf("write error: %w, close error: %v", err, closeErr)
			} else {
				err = closeErr
			}
		}
	}()

	_, err = file.Write(data)
	return err
}

// If for whatever reason, you can't handle an error, it's better to just log the error.
// Alternatively, one could choose to ignore the error, but the book suggests that
// hacing documentation as to why the error is being ignore is always best
