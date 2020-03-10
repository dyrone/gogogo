package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	fmt.Printf("init buffer len: %d\n", b.Len())
	fmt.Printf("init buffer cap: %d\n", b.Cap())
	b.Grow(64)
	fmt.Printf("buffer len after grow 64: %d\n", b.Len())
	fmt.Printf("buffer cap after grow 64: %d\n", b.Cap())
	bb := b.Bytes()
	b.Write([]byte("64 bytes or fewer"))
	// if the write bytes lenth exceed the cap of buffer, go routine will panic
	//b.Write([]byte("64 bytes or feweri dsa jdsa djskla djskl djskla djsakl djskl dj sakld jsakl dsa"))
	fmt.Printf("buffer len after write: %d\n", b.Len())
	fmt.Printf("buffer cap after write: %d\n", b.Cap())
	fmt.Printf("%q", bb[:b.Len()])
}
