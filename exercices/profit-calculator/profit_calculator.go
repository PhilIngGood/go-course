package main

import (
	"fmt"
)

func main() {

	revenues := userInput("Revenues: ")
	expenses := userInput("Expenses: ")
	taxRate := userInput("Tax Rate: ")
	//	fmt.Print("Enter your revenues: ")
	//	fmt.Scan(&revenues)
	//
	//	fmt.Print("Enter your expenses: ")
	//	fmt.Scan(&expenses)
	//
	//	fmt.Print("Tax Rate: ")
	//	fmt.Scan(&taxRate)
	//
	//	EBT := revenues - expenses
	//	fmt.Println("Raw earning: ", EBT)
	//
	//	profit := EBT * (1 - taxRate/100)
	//	ratio := EBT / profit

	ebt, profit, ratio := profitCalculations(revenues, expenses, taxRate)

	fmt.Println("Raw earning: ", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}

func userInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func profitCalculations(revenues, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenues - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
