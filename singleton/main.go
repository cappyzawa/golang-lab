package main

// Original code is https://text.baldanders.info/golang/singleton-pattern/
import (
	"fmt"
	"sync"
)

//Hello class
type Hello struct{}

func (h *Hello) String() string {
	return "Hello"
}

var instance *Hello
var once sync.Once

//GetInstance returns singleton instance
func GetInstance() *Hello {
	once.Do(func() {
		instance = &Hello{}
	})
	return instance
}

func main() {
	ch := make(chan interface{})
	go run(ch, "Alice")
	go run(ch, "Bob")
	go run(ch, "Chris")
	<-ch
	<-ch
	<-ch
}

func run(ch chan<- interface{}, person string) {
	hello := GetInstance()
	ch <- hello //blocking
	fmt.Println(hello, person)
}
