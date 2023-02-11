package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Press 1 for check Roll number\nPress 2 for check variable type")
	fmt.Scan(&n)
	if n == 1 {
		var name string
		fmt.Print("Enter your name :")
		fmt.Scan(&name)
		switch name {

		case "Umar", "umar":
			fmt.Print("Your rollnumber is : 02")
		case "Tayyab", "tayyab":
			fmt.Print("Your rollnumber is : 43")
		case "Rajab", "rajab":
			fmt.Print("Your rollnumber is : 24")
		case "Mudassar", "mudassar":
			fmt.Print("Your rollnumber is : 27")
		default:
			fmt.Print("You are out of friend group!")
		}

	} else if n == 2 {
		var n interface{} = 2 + 3i
		switch n.(type) {
		case int:
			fmt.Print("You enter a Integer variable")
		case float32:
			fmt.Print("You enter a Float variable")
		case string:
			fmt.Print("You enter a String variable")
		case complex128:
			fmt.Print("You enter a Complex variable")

		}

	} else {
		fmt.Print("You enter wrong input!!")

	}

}
