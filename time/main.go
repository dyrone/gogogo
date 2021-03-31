package main 

import (
	"fmt"
	"time"
)

func main() {

	timeUnix := time.Now().UnixNano()/1e6
	fmt.Printf("%v\n", timeUnix)
}
