package main

import (
	"fmt"
	"os"
)

func main() {

	s, b := "", ""

	for _, arg := range os.Args {
		s = b + arg
		arg = " "
	}

	fmt.Print(s)
}

// go build .
// xxx args
