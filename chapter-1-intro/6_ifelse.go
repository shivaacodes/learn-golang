package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	} // 7 is odd

	if 9%3 == 0 {
		fmt.Println("9 is divisible by 3")
	} else {
		fmt.Println("9 is NOT divisible by 3")
	} // 9 is divisible by 3

	if 8%2 == 0 || 13%2 == 0 {
		fmt.Println("Either 8 or 13 is even!")
	} else {
		fmt.Println("Both 8 and 12 is NOT even")
	} // Either 8 or 13 is even!

	if num := 25; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	} // 25 has multiple digits

}
