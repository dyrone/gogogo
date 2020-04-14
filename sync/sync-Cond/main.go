package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func listen(name string, a map[string]int, c *sync.Cond) {
	println(name, ": before lock ...")
	c.L.Lock()
	println(name, ": after lock ...")
	println(name, ": before wait...")
	c.Wait()
	fmt.Println(name, " age:", a["T"])
	c.L.Unlock()
}

func broadcast(name string, a map[string]int, c *sync.Cond) {
	time.Sleep(time.Second)
	println(name, ": before lock ...")
	c.L.Lock()
	println(name, ": after lock ...")
	a["T"] = 25
	c.Broadcast()
	c.L.Unlock()
}

func main() {

	var age = make(map[string]int)

	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	// listener 1
	go listen("lis1", age, cond)

	// listener 2
	go listen("lis2", age, cond)

	// broadcast
	go broadcast("b1", age, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

/**
	Recently, I stumbled into the use case where there will be a single goroutine retrieving a data asynchronously and multiple goroutines will wait before they can read the data retrieved by the first goroutine.
	We cannot use channels because then only one goroutine can read the data. We could use a mux locker for write operation into a global variable but then there will be additional complexity to inform the waiting goroutines to start reading the data.
	Golang provides a built-in solution for this use case in the sync package. We could use sync.Cond to make multiple goroutines to pause their execution and wait before the Cond is free again.
We will see a simple code stub in which 2 goroutines waits, using sync.Wait(), for another goroutine to broadcast using sync.Broadcast().
**/

/*
	Happens here:

	* Maps are reference types, same as slices and pointers. So, we directly passed the map age to the functions.
	* We made a chan of type os.Signal and relayed os.Interrupt signal to this channel using signal.Notify() so that we can end the program from the terminal using ctrl+c.
	* The sync.NewCond() takes a sync.Locker interface as argument and returns pointer reference to sync.Cond instance. The sync.Mutex instance satisfies the Locker interface since it implements the Lock() and Unlock() functions.
	* Basically, the time.Sleep(time.Second) is added inside the broadcast() to ensure that the cond.Broadcast() is called after the cond.Wait() functions are called first in the listen() functions which are called in separate goroutines.
	* The thing to note is that if cond.Wait() before cond.Broadcast() then the listeners will keep on waiting forever.

	From: https://medium.com/@pinkudebnath/head-first-into-sync-cond-of-golang-be71779699b1
*/
