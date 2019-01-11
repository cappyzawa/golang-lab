package main

import (
	"fmt"
	"sync"
)

func f() {
	fmt.Println("Do")
}

func main() {
	var once sync.Once
	once.Do(f)
	once.Do(f)
	twice(once)
	fmt.Println("done")
}

func twice(once sync.Once) {
	once.Do(f)
}
