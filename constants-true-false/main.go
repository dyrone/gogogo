package main

import "fmt"
import "reflect"

const (
	true        = 0 == 0 // Untyped bool.
	false       = 0 != 0 // Untyped bool.
	truee       = 0 == 0
	trueee bool = 0 == 0
)

func main() {
	fmt.Printf("untyped bool true : %v, type: %v\n", true, reflect.TypeOf(true))
	fmt.Printf("untyped bool false : %v\n", false)
	fmt.Printf("untyped bool truee: %v\n", truee)
	fmt.Printf("untyped bool trueee: %v\n", trueee)

	if true == true {
		fmt.Println("true is true!")
	}
}
