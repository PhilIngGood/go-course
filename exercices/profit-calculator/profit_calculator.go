package main

import (
	"errors"
	"fmt"
	"os"
)

// Goal
// 1) Validate user input :
// -> show error message & exit if invalid input is provided
// -> No negative number
// -> No 0
//
// 2) Store calculated results into file

const outputFile = "calculations.txt"

func main() {
	revenues, errR := userInput("Revenues: ")
	expenses, errE := userInput("Expenses: ")
	taxRate, errT := userInput("Tax Rate: ")

	if errR != nil || errE != nil || errT != nil {
		fmt.Println(errR, errE, errT)
		return
	}
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
	storeResults(ebt, profit, ratio)

	fmt.Println("Raw earning: ", ebt)
	fmt.Println("Profit:", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(outputFile, []byte(results), 0644)
}

func userInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("[ERR] This value can't be lesser than 0]")
	}

	return userInput, nil
}

func profitCalculations(revenues, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenues - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
