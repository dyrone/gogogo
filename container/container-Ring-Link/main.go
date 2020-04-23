package main

import "fmt"
import "container/ring"

func main() {
	r := ring.New(2)
	s := ring.New(2)

	for i := 0; i < r.Len(); i++ {
		r.Value = 0
		r = r.Next()
	}

	for j := 0; j < s.Len(); j++ {
		s.Value = 1
		s = s.Next()
	}

	rs := r.Link(s)

	rs.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
