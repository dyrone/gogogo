package main

import "fmt"

func main() {
	str := "aaa"

	fmt.Println("First, test value pass:" + str)
	fmt.Println("begin:" + str)
	testPassStringValue(str)
	fmt.Println("after:" + str)

	fmt.Println("Second, test pointer pass:" + str)
	fmt.Println("begin:" + str)
	testPassStringPointer(&str)
	fmt.Println("after:" + str)
}

func testPassStringPointer(str *string) {
	*str = "bbb"
	fmt.Println("in testPassStringPointer:" + *str)
}

func testPassStringValue(str string) {
	str = "bbb"
	fmt.Println("in testPassStringValue:" + str)
}
