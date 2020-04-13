package main

import "fmt"

func main() {
	b1 := []byte("hello")
	b2 := []byte{119, 0, 114, 0}

	copy(b1, b2)
	fmt.Printf("b1: %s\n", b1)
	fmt.Printf("b2: %s\n", b2)

}
