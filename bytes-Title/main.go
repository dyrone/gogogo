package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.Title([]byte("her royal highness")))
	fmt.Printf("%s\n", bytes.Title([]byte("哈哈her royal highness")))
}
