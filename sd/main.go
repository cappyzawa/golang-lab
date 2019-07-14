package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type ValidateResponse map[string]interface{}

type Client struct {
	Addr       string
	Jwt        string
	HttpClient *http.Client
}

func main() {
	client := &Client{
		Addr:       os.Getenv("SD_ADDR"),
		Jwt:        os.Getenv("SD_JWT"),
		HttpClient: http.DefaultClient,
	}

	url := fmt.Sprintf("%s/v4/validator", client.Addr)
	y, _ := ioutil.ReadFile("sd/screwdriver.yaml")
	body := fmt.Sprintf(`{"yaml": %s}`, fmt.Sprintf("%q", string(y[:])))
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Jwt))
	res, err := client.HttpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	var vr ValidateResponse
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&vr); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	enc := yaml.NewEncoder(os.Stdout)
	if err := enc.Encode(vr); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return
}
