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

func (u User) OutputUserDetails() { // (u User), the receiver, attach the func to the struct user, transforming it to a method of the received type
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

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
		// Values can be omitted and the default value will be choosen1
	}, nil
}
