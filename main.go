package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	d, ok := NewFileDiff(os.Args[2], os.Args[1])
	if !ok {
		os.Exit(2)
	}

	fmt.Print(d)
}
