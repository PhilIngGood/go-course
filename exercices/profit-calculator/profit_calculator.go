package main

import (
	"fmt"
)

func main() {
	var revenues, expenses, taxRate float64

	fmt.Print("Enter your revenues: ")
	fmt.Scan(&revenues)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenues - expenses
	fmt.Println("Raw earning: ", EBT)

	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit

	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

}
