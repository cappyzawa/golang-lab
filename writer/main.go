package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/tmp/tree1.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString("hello")
}
