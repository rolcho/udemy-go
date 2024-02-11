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
	//public field
	CreatedAt time.Time
}

type Admin struct {
	email string
	User
}

func (u User) OutputUserDetail() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// external package can use New by convention
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("missing input fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}

func NewAdmin(email, firstName, lastName, birthDate string) *Admin {
	return &Admin{
		email: email,
		User: User{
			firstName,
			lastName,
			birthDate,
			time.Now(),
		},
	}
}
