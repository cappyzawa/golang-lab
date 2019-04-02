package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Number    int    `yaml:"Number"`
	NumberStr string `yaml:"number_str"`
}

func main() {

	var data = `
Number: 1 
number_str: 2 
`

	var y Yaml
	err := yaml.Unmarshal([]byte(data), &y)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Printf("Number: %d\n", y.Number)
	fmt.Printf("number_str: %s\n", y.NumberStr)
}
