package main

import (
	"bytes"
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func main() {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	var buff bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			buff.WriteString("\x1b[102m[+")
			buff.WriteString(diff.Text)
			buff.WriteString("]\x1b[0m")
		case diffmatchpatch.DiffDelete:
			buff.WriteString("\x1b[101m[-")
			buff.WriteString(diff.Text)
			buff.WriteString("]\x1b[0m")
		case diffmatchpatch.DiffEqual:
			buff.WriteString(diff.Text)
		}
	}

	fmt.Println(buff.String())
}
