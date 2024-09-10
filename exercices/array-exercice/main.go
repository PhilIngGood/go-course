package main

import (
	"fmt"
	"time"
)

type products struct {
	title     string
	id        int
	price     float64
	createdAt time.Time
}

func main() {

	hobbyArray := [3]string{"kubernetes", "go", "music"}
	fmt.Println(hobbyArray)
	fmt.Println(hobbyArray[0])

	combinedHobbies := []string{"go", "music"}
	fmt.Println(combinedHobbies)

	sliceHobbies := hobbyArray[:2]
	sliceHobbies1 := hobbyArray[0:2]

	fmt.Println(sliceHobbies1)

	sliceHobbies = sliceHobbies[1:3]
	fmt.Println(sliceHobbies)

	courseGoals := []string{"LearnGo", "Create a Velero plugin"}
	courseGoals[1] = "Create a sftp velero plugin"
	courseGoals = append(courseGoals, "Become a kubernetes maintener")
	fmt.Println(courseGoals)

	productsArray := []products{
		{
			title:     "Book",
			id:        1,
			price:     9.99,
			createdAt: time.Now(),
		},
		{
			title:     "Chair",
			id:        2,
			price:     49.99,
			createdAt: time.Now(),
		},
	}

	fmt.Println(productsArray)

	laptop := products{
		title:     "Laptop",
		id:        3,
		price:     1899.99,
		createdAt: time.Now(),
	}

	productsArray = append(productsArray, laptop)
	fmt.Println(productsArray)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
