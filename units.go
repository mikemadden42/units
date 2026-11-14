// Package main provides tools for converting units.
package main

import (
	"flag"
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
)

// format renders a byte count in human-readable form, e.g. 1048576 -> "1M".
func format(n uint64) string {
	return bytefmt.ByteSize(n)
}

func main() {
	size := flag.Uint64("size", 0, "size in bytes")
	flag.Parse()
	if flag.NFlag() == 0 || flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println(format(*size))
}
