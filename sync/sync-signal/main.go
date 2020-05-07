package main

import "os"
import "os/signal"
import "syscall"
import "fmt"

func main() {
	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)
	//s := <-sig is no difference on blocking semantics
	<-sig
	// if signal is kill(9), no output print
	fmt.Println("received the sig!")
	//fmt.Printf("interrupt by signal %s", s)
}
