package main

import "fmt"

func main() {

	var n int

	fmt.Println("Enter size of Sliceing : ")
	fmt.Scan(&n)
	array := make([]string, n)
	fmt.Println("Ente five value of Slicing")
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	array2 := array

	array[5] = "ALi"
	array = append(array, "Tayyab")

	fmt.Println("Your Array2 is : ", array2)
	fmt.Println("Your Array is  :", array)

}
