// package main

// // Create a new struct type
// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }
// type Adress struct {
// 	PinCode int
// 	City    string
// 	Country string
// 	Email   string
// }

// type Employee struct {
// 	Person   // Embedded struct
// 	Position string
// 	Adress   // Embedded struct
// }

// func main() {
// 	// Structs are collections of fields
// 	// They can be used to group data together to form records

// 	// Create a new instance of the struct
// 	var p Person

// 	// Assign values to the fields
// 	p.FirstName = "John"
// 	p.LastName = "Doe"
// 	p.Age = 30

// 	// Print the struct
// 	println(p.FirstName, p.LastName, p.Age)

// 	// You can also create and initialize a struct in one line
// 	p2 := Person{
// 		FirstName: "Jane",
// 		LastName:  "Smith",
// 		Age:       25,
// 	}

// 	println(p2.FirstName, p2.LastName, p2.Age)

// 	// One last way to create a struct
// 	p3 := Person{"Alice", "Johnson", 28}
// 	println(p3.FirstName, p3.LastName, p3.Age)

// 	// Embedding structs means using a struct as a field in another struct
// 	e := Employee{
// 		Person: Person{
// 			FirstName: "Bob",
// 			LastName:  "Brown",
// 			Age:       35,
// 		},
// 		Position: "Manager",
// 		Adress: Adress{
// 			PinCode: 12345,
// 			City:    "New York",
// 			Country: "USA",
// 			Email:   "bob.brown@mail.com",
// 		},
// 	}
// 	println(e.Introduce())

// }

// // receiver functions for structs

// // employee introducing methods
// func (e Employee) Introduce() string {
// 	return "Hello, my name is " + e.FirstName + " " + e.LastName + ", I am a " + e.Position + " living in " + e.City + ". You can reach me at " + e.Email
// }

// // person birthday method
// func (p *Person) Birthday() {
// 	p.Age++
// }

// #############################################################################

// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	name     string
// 	lastName string
// 	age      int
// 	email    string
// }

// func main() {
// 	jim := Person{
// 		name:     "Jim",
// 		lastName: "Anderson",
// 		age:      25,
// 		email:    "jim.andy@gmail.com", // don't forge the trailng comma
// 	}

// 	jimPointer := &jim
// 	jimPointer.updateName("Jimmy")
// 	jim.print()
// }

// func (p Person) print() {
// 	fmt.Printf("%+v", p)
// }

// func (personPointer *Person) updateName(newName string) {
// 	(*personPointer).name = newName
// }

// #############################################################################

package main

import "fmt"

func main() {
	// declaring and initializing a map
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"white": "#ffffff",
	// };

	// Just declaring a map, will be initalised with nil values.
	// In this case, empty map
	// colors := make(map[string]string)

	// OR use builtin make function to initialize a map
	colors := make(map[string]string)

	// assign it later, if needed
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"

	for color, hex := range colors {
		println(color, hex)
	}

	delete(colors, "red")

	// println does not print the map keys in same order
	// as they were added. Maps are unordered collections
	// So the order of keys can be different each time
	// you run the program
	fmt.Println(colors)
}
