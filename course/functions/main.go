package main

import "fmt"

// Variadic func are func that works with any amount of parameters

func main() {
	numbers := []int{1, 10, 25}
	sum := sumup(1, 10, 25)
	sumSlice := sumup(0, numbers...)
	fmt.Println(sum)
	fmt.Println(sumSlice)
}

// Function returning the sum of all integer of a slice
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
