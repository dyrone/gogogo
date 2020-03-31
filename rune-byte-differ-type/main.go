package main

import "fmt"
import "reflect"

func main() {
	s := "你好阿里巴巴"

	// int32
	for _, v := range s {
		fmt.Printf("print rune with range : %s, type: %v\n", string(v), reflect.TypeOf(v))
	}

	s1 := "hello alibaba"

	// int32
	for _, v1 := range s1 {
		fmt.Printf("print byte with range : %v, type: %v\n", v1, reflect.TypeOf(v1))
	}

	// uint8
	fmt.Printf("print byte with index: %v, type: %v\n", s1[0], reflect.TypeOf(s1[0]))
}

// https://stackoverflow.com/questions/29255746/how-encode-rune-into-byte-using-utf8-in-golang
// https://blog.golang.org/strings
