package main

import "fmt"

// Make allow the code to be a little more efficient as it will create data strcutures allocating memory depending on the size we set

// Alias example, custom type
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Alex")

	fmt.Println(userNames)

	coursesRating := make(floatMap, 3)

	coursesRating["go"] = 4.7
	coursesRating["react"] = 4.8
	coursesRating["angular"] = 4.7
	coursesRating["java"] = 4.4 // Go will have to reallocate memory for this one, as we precise a lenght of 3 for this map

	coursesRating.output()

	// for  range userNames {
	// }

	for i, v := range userNames {
		fmt.Println(i, v)
	}

	for key, val := range coursesRating {
		fmt.Println(key, val)
	}

}
