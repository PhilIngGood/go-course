package recursion

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorialRecursive(5))

}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return number
	}
	return number * factorialRecursive(number-1)
}
