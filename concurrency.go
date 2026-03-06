package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	for i := 1;i<=5;i++ {
		fmt.Println("number:",i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters(){
	for i := 'a';i<='e';i++ {
		fmt.Println("string: ",i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main(){
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
}
