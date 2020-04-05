package main

import "fmt"

func main() {
	slice := make([]byte, 20, 20)
	slice = append([]byte("hello"), 65)
	fmt.Println(string(slice))

	slice1 := append([]byte("Hello "), "world!"...)
	fmt.Println(string(slice1))

	slice2 := append([]byte("Hello "), []byte("world!")...)
	fmt.Println(string(slice2))

}
