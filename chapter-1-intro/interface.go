package main

import "fmt"

func test_interface() {
	var i any // empty interface
	describe(i)

	i = 42
	describe(i)

	i = "checking"
	describe(i)
}

func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
