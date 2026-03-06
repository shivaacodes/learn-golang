package main

import (
	"fmt"
	"sync"
	"time"
)

// error code : concurrent map writes.

// type Counter struct {
// 	v map[string]int
// }

// func (c *Counter)Inc(key string){
// 	c.v[key]++
// }

// func (c *CounteIr)Value(key string) int {
// 	return c.v[key]
// }

// func main(){
// 	c := Counter{v:make(map[string]int)}
// 	for i:=0;i<1000;i++ {
// 		go c.Inc("somekey") // concorrent writes/mutations on map
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }

// Solution code!

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter)Inc(key string){
	c.mu.Lock()
	// Lock so only 1 goroutine at a time can acces the map c.v
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter)Value(key string) int {
	c.mu.Lock()
	// Lock so only 1 goroutine at a time can access the map c.v
	defer c.mu.Unlock()
	return c.v[key]
}

func main(){
	c := SafeCounter{v:make(map[string]int)}
	for i:=0;i<1000;i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}



