package main

import "fmt"

func main() {
	age := 23
	agePointer := &age
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
	// return *age - 18
}

// The only benefit of using pointers like so is to avoid variable to get copied multiples times
// when being passed into func. So it should only be used for very big variables and NEVER for INT.
// INTs are good, INTs are life, LOVE INTs and use them in func for god sake
