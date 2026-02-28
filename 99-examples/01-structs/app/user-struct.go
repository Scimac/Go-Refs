package main

import (
	"fmt"
	"structs/utils/converter"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
}

func main() {
	firstName := getUserData("Enter First Name: ")
	lastName := getUserData("Enter Last Name: ")
	ageStr := getUserData("Enter Age: ")
	age := converter.ConvToInt(ageStr)

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		CreatedAt: time.Now(),
	}

	user.outputUserDetails()
}

func getUserData(msg string) string {
	fmt.Print(msg)
	var input string
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	return input
}

func (u *User) outputUserDetails() {
	// this method/receiver function does not modify the struct, so pointer receiver is not necessary
	fmt.Println("\nUser Details:")
	fmt.Printf("Name: %s %s\n", (*u).FirstName, (*u).LastName) // Deferencing example, can be avoided as go allows shorthand
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Created At: %s\n", u.CreatedAt.Format(time.RFC822))
}
