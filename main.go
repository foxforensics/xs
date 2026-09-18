// Experimental strings carver.
//
// Usage:
//
//	cat FILE | xs | uniq | sort > out.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var sb strings.Builder

	flush := func() {
		if len(strings.TrimSpace(sb.String())) >= 3 {
			fmt.Println(sb.String())
		}
		sb.Reset()
	}

	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "xs:", err)
			os.Exit(1)
		}
	}()

	defer flush()

	r := bufio.NewReader(os.Stdin)

	for {
		b, err := r.ReadByte()

		if err == io.EOF {
			break
		}

		if 0x19 < b && b < 0x7F {
			sb.WriteByte(b)
		} else {
			flush()
		}
	}
}
