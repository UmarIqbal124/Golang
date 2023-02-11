package main

import "fmt"

func main() {
	var i, n, fact int
	fact = 1
	fmt.Print("Enter the number whose Factorial is requried : ")
	fmt.Scan(&n)
	for i = n; i >= 1; i-- {
		fact = fact * i
	}
	fmt.Print("The fact of ", n, " is : ", fact)

}
