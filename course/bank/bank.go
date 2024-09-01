package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to GoBank")

	for {
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)

		} else if choice == 2 {
			var depositeAmount float64
			fmt.Print("Your deposite: ")
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Wrong amount, can't be null or lesser than 0")
				return
			}

			accountBalance += depositeAmount
			fmt.Println("New blance amount:", accountBalance)

		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Wrong amount, can't be null or lesser than 0")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have on your account")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("New blance amount:", accountBalance)

		} else {
			fmt.Println("Goodbye !")
			// return
			break
		}
	}

	fmt.Println("Thanks for choosing GoBank")
}
