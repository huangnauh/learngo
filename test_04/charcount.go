package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Println("rune count")
	for c, n := range counts {
		fmt.Printf("%q %d\n",c,n)
	}
	fmt.Println("len count")
	for c, n := range utflen {
		if c > 0 {
			fmt.Printf("%d %d\n",c,n)
		}
	}
	if invalid > 0 {
		fmt.Printf("%d invalid chars\n", invalid)
	}
}
