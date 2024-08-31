package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investAmount, years, expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investAmount)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Output informations
	// fmt.Println("Future value:", futureValue)
	// fmt.Println("Future value considering inflation:", futureRealValue)
	fmt.Printf("Future value: %.2f\nFuture value considering inflation: %.2f\n", futureValue, futureRealValue)

}
