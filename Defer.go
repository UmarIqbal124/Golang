package main

import "fmt"

func main() {
	defer fmt.Println("You")
	fmt.Println("I")
	fmt.Println("Love")
}
