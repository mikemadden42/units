package main

import (
	"flag"
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
)

func main() {
	size := flag.Uint64("size", 0, "size in bytes")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println(bytefmt.ByteSize(*size))
}
