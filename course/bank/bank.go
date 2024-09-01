package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceString), 0644)
}

func getBalanceFromFile() (float64, error) {
	dataByte, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(dataByte)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance values")
	}

	return balance, nil
}

func main() {
	// var accountBalance float64 = os.ReadFile(accountBalanceFile)
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

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
			writeBalanceToFile(accountBalance)

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
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye !")
			fmt.Println("Thanks for choosing GoBank")
			return
		}
	}
}
