package main

import "fmt"

func main() {
	// Dynamic array
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[0:1])

	udatedPrices := append(prices, 5.98)
	fmt.Println(udatedPrices)
}

// func main() {
// 	// Array declaration
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	// Set a value
// 	productNames[2] = "A carpet"

// 	fmt.Printf("prices values at the start: %v\n", prices)
// 	fmt.Printf("product names: %v\n", productNames)

// 	// Read value from array
// 	fmt.Printf("3rd element of prices equals: %v\n", prices[2])

// 	fmt.Println("Slicing the prices array")
// 	// Slices
// 	featuredPrices := prices[1:4]
// 	// featuredPrices := prices[:3]
// 	// featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:2]

// 	fmt.Printf("featurePrices value: %v\n", featuredPrices)
// 	fmt.Printf("highlightedPrices values: %v\n", highlightedPrices)

// 	fmt.Println("Changing 3rd elements of prices & 1rst element of featuredPrices")
// 	prices[2] = 99.9
// 	featuredPrices[0] = 157.66

// 	fmt.Printf("prices values at the end: %v\n", prices)
// 	fmt.Printf("feature prices: %v\n", featuredPrices)
// 	fmt.Printf("highlightedPrices values: %v\n", highlightedPrices)

// 	highlightedPrices = featuredPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	// We can reslice a slice to increase it's lenght, but only toward the end of the initial array, like shown in the 3 lines above.
// 	// highlightedPrices has ben resized using itself toward the end of the initial slice

// }
