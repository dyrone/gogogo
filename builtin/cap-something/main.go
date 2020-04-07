package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("array cap: %d\n", cap(array))
	fmt.Printf("array len: %d\n", len(array))
}
