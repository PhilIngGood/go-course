package main

import (
	"fmt"

	"course.go/structs/user"
)

func main() {
	userFirtName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	// ... do something, anything will do

	var appUser *user.User
	appUser, err := user.New(userFirtName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
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

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
