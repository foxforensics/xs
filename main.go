// go-template.
//
// Usage:
//
//	go-template arg
//
// The arguments are:
//
//	arg
//	    Argument (required).
package main

import (
	"fmt"
	"os"
)

var Usage = `© 2026 Fox Forensics. Licensed under MIT License.
Usage: go-template ARG

Report bugs at: foxforensics.eu/issues
`

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprint(os.Stderr, Usage)
		os.Exit(2)
	}
}
