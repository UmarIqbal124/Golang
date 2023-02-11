package main

import "fmt"

func main() {

	/*

		*
		**
		***
			var i, j, n int
			fmt.Print("Enter the number of line : ")
			fmt.Scan(&n)

			for i = 1; i <= n; i++ {
				for j = 1; j <= i; j++ {
					fmt.Print("* ")
				}
				fmt.Printf("\n")
			}

		***
		**
		*

	        var i, j, n int
		    fmt.Print("Enter the number of line : ")
	    	fmt.Scan(&n)

	    	for i = n; i >= 1; i-- {

		     	for j = 1; j <= i; j++ {
			    	fmt.Print("* ")
		     	}
		    	fmt.Printf("\n")
	    	}

	      *
		 **
		***

	    	var i, j, n, k int
	    	fmt.Print("Enter the number of line : ")
	    	fmt.Scan(&n)

			for k = 1; k <= n; k++ {

				for i = n - 1; i >= k; i-- {
					fmt.Print(" ")
				}

				for j = 1; j <= k; j++ {
					fmt.Print("*")
				}
				fmt.Printf("\n")
			}

	    ***
		**
		*

			var i, j, n, k int
			fmt.Print("Enter the number of line : ")
			fmt.Scan(&n)

			for k = n; k >= 1; k-- {

				for i = n - 1; i >= k; i-- {
					fmt.Print(" ")
				}

				for j = 1; j <= k; j++ {
					fmt.Print("*")
				}
				fmt.Printf("\n")
			}

	*/

	// Dimond Shape
	var i, j, n, k int
	fmt.Print("Enter the number of line : ")
	fmt.Scan(&n)

	for k = 1; k <= n; k++ {

		for i = n - 1; i >= k; i-- {
			fmt.Print(" ")
		}

		for j = 1; j <= k; j++ {
			fmt.Print("* ")
		}
		fmt.Printf("\n")
	}
	for k = n - 1; k >= 1; k-- {

		for i = n - 1; i >= k; i-- {
			fmt.Print(" ")
		}

		for j = 1; j <= k; j++ {
			fmt.Print("* ")
		}
		fmt.Printf("\n")
	}

}
