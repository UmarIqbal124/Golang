package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now().Weekday()
	switch time.Saturday {
	case currentTime + 0:
		fmt.Println("Today is Saturday")
	case currentTime + 1:
		fmt.Println("Today is Friday")
	case currentTime + 2:
		fmt.Println("Today is Thusday")
	case currentTime + 3:
		fmt.Println("Today is Wednasday")
	case currentTime + 4:
		fmt.Println("Today is Tuesday")
	case currentTime + 5:
		fmt.Println("Today is Monday")
	default:
		fmt.Println("Today is Sunday")

	}
}
