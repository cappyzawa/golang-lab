package main

import (
	"fmt"
	"time"
)

// Checker checks status periodically.
// type Checker struct {
// Service  reminder.LineService
// Duration time.Duration
// }

// Check checks targets's status
// func (c *Checker) Check(targets []string) error {
func main() {
	ticker := time.NewTicker(1 * time.Second).C

	for {
		select {
		case <-ticker:
			// put logics here
			fmt.Println("foo")
		}
	}
}
