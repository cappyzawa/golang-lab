package main

import (
	"errors"
	"log"
	"net/http"
)

type retry struct {
	max       int
	transport http.RoundTripper
}

func (r *retry) RoundTrip(req *http.Request) (*http.Response, error) {
	for i := 0; i < r.max; i++ {
		log.Println("Attemt: ", i+1)
		res, err := r.transport.RoundTrip(req)
		if res != nil && err == nil {
			return res, err
		}
		log.Println("Retring...")
	}
	return nil, errors.New("retry max exceed")
}

func main() {
	r := &retry{
		max:       3,
		transport: http.DefaultTransport,
	}
	c := &http.Client{Transport: r}
	req, _ := http.NewRequest(http.MethodGet, "https://google.co1m", nil)
	res, err := c.Do(req)
	if err != nil {
		log.Fatalln("failed to do request: ", err)
	}
	log.Println("status: ", res.StatusCode)
}
