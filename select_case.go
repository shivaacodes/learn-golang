package main

import (
	"fmt"
	"time"
)

func main(){
	tick := time.Tick(100 * time.Millisecond) // channel by default
	boom := time.After(500 * time.Millisecond)

	for { // infinite loop until a return OR break
		select{ 
		case <-tick: // if recieve anything from tick channel
		  fmt.Println("tick")
		case <-boom: // if recieve anything from boom channel
			fmt.Printf("BOOM!")
			return // exit the infinite loop
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//Output

//     .
//     .
//tick
//     .
//     .
//tick
//     .
//     .
//tick
//     .
//     .
//tick
//     .
//     .
//BOOM!
