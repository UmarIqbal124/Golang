package main

import "fmt"

func main() {
	var i, j int
	fmt.Print("Enter 2 number to perform mathematic oprations : ")
	fmt.Scan(&i, &j)
	calculation(i, j)
	value, value1, value2, value3 := calculation(i, j)

	fmt.Println("The Sum of given numbers is : ", value)
	fmt.Println("The Sub of given numbers is : ", value1)
	fmt.Println("The Mult of given numbers is : ", value2)
	fmt.Println("The Dev of given numbers is : ", value3)
}

func calculation(x int, y int) (int, string, int, int) {
	Sum := x + y
	Mult := x * y
	Dev := x / y
	name := "Umar"

	return Sum, name, Mult, Dev
}
