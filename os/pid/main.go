package main

import "fmt"
import "os"

func main() {
	pid := os.Getpid()
	fmt.Printf("pid: %d", pid)
}
