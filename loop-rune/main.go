package main

import "fmt"

func main() {
	s := "你好阿里巴巴"

	for _, v := range s {
		fmt.Printf("%v\n", v)
		fmt.Printf("%s\n", string(v))
	}

	s1 := "hello alibaba"

	for _, v1 := range s1 {
		fmt.Printf("%v\n", v1)
	}

	s2 := "截取中文"
	//试试这样能不能截取?
	res := []rune(s2)
	fmt.Println(string(res[:2]))
}
