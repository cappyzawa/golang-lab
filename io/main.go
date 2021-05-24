package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type cli struct {
	out io.Writer
	err io.Writer
	in  io.Reader
}

func (c *cli) Tee(args []string) int {
	if len(args) <= 1 {
		fmt.Fprintf(c.err, "required an arg for filename\n")
		return 1
	}
	filename := args[1]
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(c.err, "create a file: %v\n", err)
		return 2
	}
	defer f.Close()

	tee := io.TeeReader(c.in, f)
	s := bufio.NewScanner(tee)
	for s.Scan() {
		fmt.Fprintf(c.out, "%s\n", s.Text())
	}
	return 0
}

func main() {
	c := &cli{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	os.Exit(c.Tee(os.Args))
}
