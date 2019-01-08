package main

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	l, err := net.Listen("tcp", ":6060")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on", l.Addr())
	go http.Serve(l, nil)
	for {
		fib(30)
	}
}
