package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	wg.Add(1)
	go say("Hi")
	wg.Add(1)
	go say("Bye")
	wg.Wait()
}
