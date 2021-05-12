package main

import (
	"fmt"

	"github.com/gookit/color"
)

func main() {
	c := color.FgColors["red"].Sprint("ffffffff")
	fmt.Println(c)
}
