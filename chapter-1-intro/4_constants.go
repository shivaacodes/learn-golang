package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants() {
	fmt.Println(s)

	const n = 76652662

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
