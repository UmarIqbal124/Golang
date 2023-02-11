package main

import "fmt"

func main() {
	var i, j, n int
	fmt.Print("Enter the number of line : ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Println("* ")
		}
	}
}
