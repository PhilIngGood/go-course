package functionarevalues

import "fmt"

// Functions in go can be used as value in functions.
// Here, the multiplyNumbers function accepts an other functions as a parameter. That allow us to write more flexible code
// Functions can also return functions

// An alias as been set to this function type
type multiplyFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := multiplyNumbers(&numbers, double)
	tripled := multiplyNumbers(&numbers, triple)
	fmt.Printf("Original slice :%v\n", numbers)
	fmt.Printf("Doubled slice :%v\n", doubled)
	fmt.Printf("Tripled slice :%v\n", tripled)
}

// Func returning function
func getTransformerFunction() multiplyFn {
	return double
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func multiplyNumbers(numbers *[]int, multiply multiplyFn) []int {
	multipliedNumber := make([]int, len(*numbers))

	for index, value := range *numbers {
		multipliedNumber[index] = multiply(value)
	}

	return multipliedNumber
}
