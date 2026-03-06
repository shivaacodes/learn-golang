package main

import "fmt"

func sendData(ch chan string){
	ch <- "hello from goroutine"
}

func main(){
	ch := make(chan string)

	go sendData(ch)

	msg := <-ch
	fmt.Println(msg)
}
