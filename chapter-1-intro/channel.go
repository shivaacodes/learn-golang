// channels are used to synchronize/ orchestrate/ communicate between goroutines

package main

import "fmt"

// channels: is like a tunnel which transfers data between goroutines.
// it is a container datatype. we have to specify the type of the data.

func sum(s []int, c chan int) { // here c is a channel which take int datatype
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sending teh result to teh channel (here; sending sum to c)
}

// 5 ref types : slices, maps, funcs, channels, pointers
func main() {
	s := []int{3, 4, 6, -8, 9, 2}

	c := make(chan int) // ref types can be declared with a make syntax

	// spawn 2 goroutine
	go sum(s[:len(s)/2], c) // sum of first half
	go sum(s[len(s)/2:], c) // sum of first half

	x, y := <-c, <-c // receive data from c (blocking execution)

	fmt.Println(x, y, x+y)

}

// 2 types of channels (high-level difference)

// buffered : queue; goroutines can push values into a queue, sequentially until the size is filled.
// (keep pushing it wont block)

// unbuffered : no queue; just a tunnel. until the producer and the consumer (goroutines) are in sync.
// until the recieving goroutine take the value; it is blocked. usecase: for forced synchronization
