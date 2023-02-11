package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter a number whose Table is requried : ")
	fmt.Scan(&n)
	fmt.Println("****** Table ******")
	for i := 1; i <= 10; i++ {
		j := n * i
		fmt.Println(n, " X ", i, " = ", j)
	}

}
