package main

import "fmt"

type Student struct {
	StdName, StdNumber string
	StdId              int
}

func main() {

	/* Deceleration and Initialization
	Std := Student{
		StdId:     03,
		StdName:   "Umar G",
		StdNumber: "03475272178",
	}

	fmt.Println(Std)
	*/
	var std, std2 Student

	fmt.Print("Enter student Id :")
	fmt.Scan(&std.StdId)
	fmt.Print("Enter student name :")
	fmt.Scan(&std.StdName)
	fmt.Print("Enter student number :")
	fmt.Scan(&std.StdNumber)

	std2 = std

	std.StdName = "Tayyab"
	fmt.Println("\n", std2)
	fmt.Println("\n", std)

}
