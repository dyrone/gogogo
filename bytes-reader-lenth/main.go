package main

import (
	"bytes"
	"fmt"
)

func main() {
	reader := bytes.NewReader([]byte("Hi!"))
	fmt.Printf("reader init len: %d\n", reader.Len())
	//fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len())
	b, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("read a byte: %v\n", b)
	fmt.Printf("reader len after read a byte: %d\n", reader.Len())
}
