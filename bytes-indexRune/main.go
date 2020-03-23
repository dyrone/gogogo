package main

import "fmt"
import "bytes"

func main() {
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))
	fmt.Println(bytes.IndexRune([]byte("你好，世界"), '世'))
}
