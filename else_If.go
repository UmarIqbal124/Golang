package main

import (
	"fmt"
)

func main() {
	var marks int
	fmt.Print("Enter your marks : ")
	fmt.Scan(&marks)
	if marks >= 85 {
		fmt.Print("Your grade is : A+")
	} else if marks >= 80 {
		fmt.Print("Your grade is : A")
	} else if marks >= 75 {
		fmt.Print("Your grade is : B+")
	} else if marks >= 70 {
		fmt.Print("Your grade is : B")
	} else if marks >= 65 {
		fmt.Print("Your grade is : C+")
	} else if marks >= 60 {
		fmt.Print("Your grade is : C")
	} else if marks >= 55 {
		fmt.Print("Your grade is : D")
	} else if marks >= 50 {
		fmt.Print("Your grade is : E")
	} else {
		fmt.Print("Your grade is : F")
	}
}
