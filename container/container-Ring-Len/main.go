package main

import "fmt"
import "container/ring"

func main() {
	r := ring.New(4)

	fmt.Println(r.Len())
}
