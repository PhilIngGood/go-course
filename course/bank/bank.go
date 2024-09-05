package main

import (
	"fmt"

	"course.go/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	// var accountBalance float64 = os.ReadFile(accountBalanceFile)
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to GoBank")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositeAmount float64
			fmt.Print("Your deposite: ")
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Wrong amount, can't be null or lesser than 0")
				continue
			}
			accountBalance += depositeAmount
			fmt.Println("New blance amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			var withdrawAmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Wrong amount, can't be null or lesser than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have on your account")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("New blance amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye !")
			fmt.Println("Thanks for choosing GoBank")
			return
		}
	}
}
