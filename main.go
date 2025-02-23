package main

import (
	"flag"
	"fmt"
	"os"
)

var sd = flag.Bool("s", false, "string diff")

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "usage: diff [-s] aaa bbb")
		os.Exit(1)
	}

	switch {
	case *sd:
		d := NewStringDiff(flag.Arg(1), flag.Arg(0))
		fmt.Print(d)
	default:
		d, ok := NewFileDiff(flag.Arg(1), flag.Arg(0))
		if !ok {
			os.Exit(2)
		}
		fmt.Print(d)
	}
}
