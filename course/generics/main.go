package main

import "fmt"

func main() {
	a := 6
	b := 5
	checkValueType(a)
	res := add(a, b)
	fmt.Println(res)
}

// Example of func accepting any type of value
func checkValueType(value interface{}) {
	// Do something
	typedVal, ok := value.(int)
	if ok {
		fmt.Printf("Value %v is an integer\n", typedVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ")
	// case float64:
	// 	fmt.Println("Float64: ")
	// case string:
	// 	fmt.Println("String: ")
	// default:

	// }

}

// Generic example
func add[T int | float64 | float32 | string](a, b T) T {
	return a + b
}
