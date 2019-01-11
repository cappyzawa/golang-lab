package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var questions = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func main() {
	var count int
	inputCh := input(os.Stdin)
	timer := time.After(5 * time.Second)
	for _, q := range questions {
		fmt.Println(q)
		select {
		case answer := <-inputCh:
			if answer == q {
				count++
			}
		case <-timer:
			fmt.Println()
			fmt.Println("time over")
			fmt.Println("your correct answers", count)
			os.Exit(0)
		}
	}
}
