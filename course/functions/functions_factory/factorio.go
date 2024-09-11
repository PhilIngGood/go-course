package function_factory

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	doubleFn := transformFactory(2)
	tripleFn := transformFactory(3)

	doubled := transformNumbers(&numbers, doubleFn)
	tripled := transformNumbers(&numbers, tripleFn)

	// functions defined in parameter is an anonymous function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(numbers)
	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// This function is called a factory as it is returning a function
func transformFactory(factor int) func(number int) int {
	return func(number int) int {
		return number * factor
	}
}
