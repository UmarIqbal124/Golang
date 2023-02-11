package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int

	fmt.Println("\n****************************************************")
	fmt.Println("\tWellCome in friend Casino")
	fmt.Printf("*****************************************************\n\n")

	fmt.Printf("\nPress 1 for Play Casino\n")
	fmt.Printf("Press 2 for other services\n")
	fmt.Print("\tPress : ")
	fmt.Scan(&num)

	if num == 1 {
		Play()

	} else {
		fmt.Println("\nWellCome for joining us")
	}

}

func Play() {
	var firstName, secondName string
	var totalAmount, dataBase, betAmount, luckyNumber, random, check int
	fmt.Print("\nPlease Enter Your First Name (necessory): ")
	fmt.Scan(&firstName)
	fmt.Print("Please Enter Your Second Name (optional): ")
	fmt.Scan(&secondName)

	fmt.Println("\n\nWellCome Mr/Miss ", firstName+" "+secondName, " for joining us")

	fmt.Print("\nPlease Enter Your Total amount : ")
	fmt.Scan(&totalAmount)
	dataBase = totalAmount

	for check != 2 {

		fmt.Print("Please Enter Your Bet amount ")
		fmt.Scan(&betAmount)
		fmt.Print("Please Enter your lucky numnber Between 0 - 30 : ")
		fmt.Scan(&luckyNumber)

		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(30)

		if luckyNumber == random {
			fmt.Print("\n\nCongratulation you win successfully")
			totalAmount = totalAmount + (betAmount * 10)
			fmt.Print("\nNow your total amount is : ", totalAmount)

		} else {
			fmt.Print("\n\nSorry you loss the game\nThe lucky number was : ", random)
			totalAmount = totalAmount - betAmount
			fmt.Print("\nNow your total amount is : ", totalAmount)
		}

		fmt.Print("\n\nPress 1 for play again \nPress 2 for exit \n    Press :")
		fmt.Scan(&check)
	}

	fmt.Print("\n\n\nThanks for coming \nWhen you come your total amount was ", dataBase, "\n Now your total amount is : ", totalAmount)
	if totalAmount > dataBase {
		fmt.Print("\nHave a good day\n\n\n")
		time.Sleep(10 * time.Second)
	} else {
		fmt.Print("\nBetter lucK for next time\n")
		time.Sleep(10 * time.Second)
	}
	main()
}
