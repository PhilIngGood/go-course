package main

import "fmt"

func main() {
	// Array declaration
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	// Set a value
	productNames[2] = "A carpet"

	fmt.Println(prices)
	fmt.Println(productNames)

	// Read value from array
	fmt.Println(prices[2])

}
