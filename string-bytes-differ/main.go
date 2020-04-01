package main

import "fmt"
import "reflect"

func main() {
	bs := []byte("abc")
	fmt.Printf("type of a byte[] on index `0`: %v\n", reflect.TypeOf(bs[0]))

	//To expand, a "string" variable is fundamentally a modified slice header. The string header itself is at some location in memory (0x1040c128 in your example), and that location in memory holds two values, a pointer to an underlying array and the size of the string (unlike slices, strings don't have a "capacity" value, because, being immutable, their capacity always equals their length). When you assign to a string variable, you're changing the value of the pointer stored within that header. The address of the header itself does not change
	ss := string(bs)
	fmt.Printf("type of a string on index `0`: %v\n", reflect.TypeOf(ss[0]))
	fmt.Printf("ss address: %v\n", &ss)
	fmt.Printf("ss value: %v\n", ss)

	ss = string("def")
	fmt.Printf("ss address: %v\n", &ss)
	fmt.Printf("ss value: %v\n", ss)

}
