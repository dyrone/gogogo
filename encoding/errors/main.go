package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

type Dyrone interface {
	Compute() string
}

type MyDyrone struct {
	name string
	age  int
}

func (d MyDyrone) Compute() string {

	return fmt.Sprintf("Dyrone name: %s, age: %d\n", d.name, d.age)
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}

	myDyrone := MyDyrone{"Teng Long", 32}
	fmt.Printf("MyDyrone : %#v\n", myDyrone)
}
