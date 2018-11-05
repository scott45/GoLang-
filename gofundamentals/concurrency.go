package main

import (
	"fmt"
	"time"
	"sync"
	// "runtime"
)

func main () {
	// runtime.GOMAXPROCS(2)
	var WaitGrp sync.WaitGroup
	WaitGrp.Add(2)
	
	go func() {
	defer WaitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("heyy man")
	}()

	go func() {
		defer WaitGrp.Done()

		fmt.Println("heyy woman")
	}()

	WaitGrp.Wait()
}

// channels
// myChannel := make(chan int)
// myChannel := make(chan int, 5)