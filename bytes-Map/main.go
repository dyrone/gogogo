package main

import (
	"bytes"
	"fmt"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+14)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+14)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}
