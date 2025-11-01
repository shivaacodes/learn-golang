package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") //goroutine: concurrency
	say("hello")
}

// Output

// hello
// world
// world
// hello
// hello
// world
// world
// hello
// hello
