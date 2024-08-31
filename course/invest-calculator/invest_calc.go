package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investAmount, years, expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investAmount)

	fmt.Print("Number of years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investAmount, expectedReturnRate, years)

	// futureValue := investAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Format strings and store them
	// formatedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formatedFRV := fmt.Sprintf("Future value considering inflation: %.2f\n", futureRealValue)
	// fmt.Print(formatedFV, formatedFRV)
	// Output informations
	// fmt.Println("Future value:", futureValue)
	// fmt.Println("Future value considering inflation:", futureRealValue)
	// Using printf to format and print values 	https://pkg.go.dev/fmt@go1.23.0#Printf
	// fmt.Printf("Future value: %.2f\nFuture value considering inflation: %.2f\n", futureValue, futureRealValue)

	// Backtick string example
	fmt.Printf(`Future value: %.2f
Future value considering inflation: %.2f`, futureValue, futureRealValue)

}

// fmt.print wrapper, just as an example
func uselessFunc(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValueResult := investAmount * math.Pow(1+expectedReturnRate/100, years)
	realFuturevalue := futureValueResult / math.Pow(1+inflationRate/100, years)

	return futureValueResult, realFuturevalue
}

// Alternative way to create return values
func calculateFutureValues_bis(investAmount, expectedReturnRate, years float64) (futureValueResult float64, realFuturevalue float64) {
	futureValueResult = investAmount * math.Pow(1+expectedReturnRate/100, years)
	realFuturevalue = futureValueResult / math.Pow(1+inflationRate/100, years)

	//return futureValueResult, realFuturevalue
	return
}
