package main

import "math/rand"

//import "os"
import "fmt"

func main() {
	//rand.Seed(int64(os.Getpid()))
	rand.Seed(1)
	fmt.Printf("random int : %d\n", rand.Int63n(500))
	fmt.Printf("random int : %d\n", rand.Int63n(500))
	fmt.Printf("random int : %d\n", rand.Int63n(500))
	fmt.Printf("random int : %d\n", rand.Int63n(500))
}
