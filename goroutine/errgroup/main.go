package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	var eg errgroup.Group
	eg.Go(func() error {
		fmt.Println("Hello")
		// return errors.New("hello error")
		return nil
	})
	eg.Go(func() error {
		fmt.Println("World")
		// return errors.New("world error")
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
