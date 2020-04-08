package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("array cap: %d\n", cap(array))
	fmt.Printf("array len: %d\n", len(array))
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice cap: %d\n", cap(slice))
	fmt.Printf("slice len: %d\n", len(slice))
}
