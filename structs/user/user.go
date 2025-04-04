package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email string, password string) (*Admin, error) {
	if len(email) == 0 || len(password) == 0 {
		return nil, errors.New("please enter all fields")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil
}

func (usr User) OutputUserDetails() {
	fmt.Println(usr.firstName, usr.lastName, usr.birthdate)
}

func (usr *User) CleanName() {
	usr.firstName = ""
	usr.lastName = ""
}

func NewUser(firstName string, lastName string, birthDate string) (*User, error) {
	if len(firstName) == 0 || len(lastName) == 0 || len(birthDate) == 0 {
		return nil, errors.New("please enter all fields")
	}

	var myStr str
	myStr = "Hello world!"
	myStr.log()

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}

type str string

func (text str) log() {
	fmt.Println(text)
}
