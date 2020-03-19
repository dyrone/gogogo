package main

import "fmt"
import "bytes"

func main() {
	fmt.Printf("Fields are : %q", bytes.Fields([]byte("  foo bar  baz  ")))
}
