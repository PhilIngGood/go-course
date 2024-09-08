package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirtName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	// ... do something, anything will do
	var appUser User = User{
		firstName: userFirtName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
		// Values can be omitted and the default value will be choosen1
	}

	// This notation can also be used, to shorten the writting time
	// appUser = User{
	// 	userFirtName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// Empty structs can also be called
	// appUser = user{}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
