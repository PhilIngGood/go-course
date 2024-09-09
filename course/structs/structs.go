package main

import (
	"fmt"

	"course.go/structs/user"
)

// Alias to  string type
type str string

// Method to custom string
func (text str) log() {
	fmt.Println(text)
}

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

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	// Custom string type example
	var name str = "Phil"
	name.log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
