package main

import "fmt"

func main() {
	var u uint8
	var v = -1
	u = uint8(v)
	fmt.Printf("%T, %v", u, u)
}
