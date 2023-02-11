package main

import "fmt"

func main() {
	/*
			M := make(map[string]int) //Declaration
			M = map[string]int{
				"Umar":     02,
				"Rajab":    24,
				"Mudassar": 27,
				"Tayyab":   43,
			} //initialization


		M := map[string]int{ //Deceliration & initialization on one line
			"Umar":     02,
			"Rajab":    24,
			"Mudassar": 27,
			"Tayyab":   43,
		}

		m2 := M

		M["Ali"] = 70 //add new item
		fmt.Println(m2)
		fmt.Println(M)

		delete(m2, "Rajab")

		fmt.Println(m2)
		fmt.Println(M)                             */

	var Name, key string
	var value, n, i int

	Employee := make(map[string]int)

	fmt.Print("Enter the size of Map : ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Println("Enter the name and selery of ", i, "number employee(e.g Abubakr 2400) : ")
		fmt.Scan(&Name, &value)
		Employee[Name] = value

	}

	fmt.Print("\nYour given map is : ", Employee)
	fmt.Print("\n\n\nType a key to search in map : ")
	fmt.Scan(&key)

	Value, check := Employee[key]
	if check == true {
		fmt.Print("\n\nYes your given key is in map and its value is : ", Value)
	} else {
		fmt.Print("\n\nSorry your given key is not in map\n")
	}
}
