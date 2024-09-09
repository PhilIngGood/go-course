package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// This notation can also be used, to shorten the writting time
// appUser = User{
// 	userFirtName,
// 	userLastName,
// 	userBirthDate,
// 	time.Now(),
// }
//
// Empty structs can also be called
// appUser = user{}

// Embeded struct (Java inheritence)
type Admin struct {
	email    string
	password string
	User
}

//  Here, User is anonimized as we did not set any name to this object. That allow us to call User method for Admin object directly
// For example :
// if non anonimized, we have to use admin.user.OutputUserDetails()
// if anonimized, we can do admin.OutputUserDetails

// (u User), the receiver, attach the func to the struct user, transforming it to a method of the received type
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Method
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""

}

// User struct builder
func New(firstName, lastName, birthDate string) (*User, error) {
	// validations
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("wrong Users parameters")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
		// Values can be omitted and the object default value will be choosen1
	}, nil
}

// Admin builder
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}
