package main

import "fmt"

//Here, below bothe fns are same which is bad so we use generics

// slice of ints

// func IndexInt(s []int, x int) int {
// 	for i, v := range s {
// 		if v == x {
// 			return i
// 		}
// 	}
// 	return -1
// }

// slice of strings

// func IndexString(s []string, x string) int {
// 	for i, v := range s {
// 		if v == x {
// 			return i
// 		}
// 	}
// 	return -1
// }

// generics example : tl;dr we use types as parameters (reduce boilerplate code)

func Index[T comparable](s []T, x T) int { // parameter : T , constraint : comparable
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{3, 5, 6, 8, 9}
	fmt.Println(Index(si, 8))

	ss := []string{"get", "push", "pull"}
	fmt.Println(Index(ss, "pull"))
}
