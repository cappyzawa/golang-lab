package main

import (
	"fmt"
	"github.com/cappyzawa/golang-lab/cobra/cmd"
	"os"
)

func main() {
	command := cmd.NewCmd()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
