package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var pointer *float64
	fmt.Print("Enter a number whose Saqure root in requried : ")
	fmt.Scan(&n)

	b := math.Sqrt(float64(n))

	pointer = &b

	fmt.Print("\n\nThe square root of ", n, " is ", *pointer)
}
