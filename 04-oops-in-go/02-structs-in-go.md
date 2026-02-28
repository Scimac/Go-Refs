### Structs Basics

- Structs are collection of named elements, called fields, each of which has a name and a type.
- Useful for grouping related data.

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p1 := Person{"Alice", 30} // struct literal, positional
    
    p2 := Person{Name: "Bob", Age: 25} // named fields    
    
    var p Person // Declare, Assign values to the fields later    
	p.FirstName = "John"
	p.LastName = "Doe"
	p.Age = 30
    
    fmt.Println(person{name: "Fred"}) // Omitted fields will be zero-valued.
    fmt.Println(p1, p2)
}
```

- Struct literals, used to create structs (p1, p2), can use **named fields** or **positional fields**.
- Omitted fields will be populated with zero-values.
### Constructor or Factory functions

- Functions used to create the structs. These are not in-built function.
- It is common practice to start the constructor function name with `new/New`.
- It is better to return the pointer to the actual value instead of a struct as it avoid creation of copy of the struct.
- Constructor functions are good way to add validations.

```go
func NewPerson(name string, age int) *Person {
	// we can add some validations here, return actual proper errors
	if firstName === '' || age <= 0  {
		fmt.Println("Invalid Data Found!")
		return nil
	}
	
    return &Person{Name: name, Age: age}
}

p := NewPerson("Alice", 30)
```

### Accessing & Updating Struct Fields

```go
p := Person{Name: "Charlie", Age: 20}
fmt.Println(p.Name) // access
p.Age = 21          // update
```

- Structs can be large, so often passed by pointer.
- Go allows shorthand field access - 

```go
p := &Person{Name: "Dana", Age: 40}
p.Age++               // same as (*p).Age++
fmt.Println(p.Age)    // 41
```

### Receiver Functions

- Functions attached to a struct type
- There is no explicit declaration of all receiver functions linked to a certain struct at one place. We need to refer to the function declaration for checking the link.

```go
func (p Person) Greet() {
    fmt.Printf("Hi, my name is %s\n", p.Name)
}

p := Person{"Eve", 28}
p.Greet()
```

### Pointer vs Value Receivers

- **Value Receiver**: 
	- Receiver functions that gets value of the struct passed as a argument.
	- It creates a copy of the struct (Behaviour of Go, as it is a **pass by value** language).
	- Does not modify original value, if the value in argument is updated. We need to return the updated value and assign it to original for change to take place, 
- **Pointer Receiver**:
	- Receiver functions that receive a pointer of the struct as an argument.
	- It can modify original, if the value is updated by dereferencing the pointer.

```go
func (p Person) IncrementAge() {
    age := p.Age++ // copy, doesn’t affect caller
    return age
}

func (p *Person) GrowUp() {
    p.Age++ // modifies original
}
```

```go
p := Person{"Frank", 18}
p.IncrementAge()
fmt.Println(p.Age) // 18 (unchanged)
p.Age = p.IncrementAge()
fmt.Println(p.Age) // 19 (modified after assignment)

p.GrowUp()
fmt.Println(p.Age) // 20 (modified directly)
```

#### Rule of thumb

- Use **pointer receivers** if:    
    - The method needs to modify the struct.        
    - The struct is large (avoid copying).        
- Use **value receivers** if:    
    - Struct is small and immutable.

### Anonymous Fields (Struct Embedding)

- Go doesn’t have inheritance, but supports **composition** via embedding.

```go
type Address struct {
    City, Country string
}

type Employee struct {
    Name string
    Address // embedded
}

func main() {
    e := Employee{
        Name:    "Grace",
        Address: Address{"London", "UK"},
    }
    fmt.Println(e.City) // directly accessible
}
```

### Struct Tags

- Add metadata (often used in JSON, DB mapping, serialization/deserialization.).

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

- Use Case - 

```go
u := User{1, "Henry"}
jsonData, _ := json.Marshal(u)
fmt.Println(string(jsonData)) // {"id":1,"name":"Henry"}
```

> In Go, `json.Marshal` is a function from the `encoding/json` package used to convert Go data structures (like structs, maps, slices, and basic types) into their JSON representation. This process is known as marshalling.

#### Basic use 

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p := Person{Name: "Alice", Age: 30, City: "New York"}
	
	// Marshal the struct into JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	
	fmt.Println(string(jsonData)) 
	// Output: {"Name":"Alice","Age":30,"City":"New York"}
}```

- The `json.Marshal` function takes a value of any type and returns a byte slice containing the JSON encoding and an error if the marshalling fails.
#### Using JSON Tags for Customisation

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int    `json:"product_id"`
	Name  string `json:"productName"`
	Price float64 `json:"price,omitempty"` // omitempty excludes if zero value
	Description string `json:"-"`           // - ignores the field
}

func main() {
	p1 := Product{ID: 1, Name: "Laptop", Price: 1200.0, Description: "Powerful laptop"}
	
	jsonData1, _ := json.Marshal(p1)
	fmt.Println(string(jsonData1)) 
	// Output: {"product_id":1,"productName":"Laptop","price":1200}
	
	p2 := Product{ID: 2, Name: "Mouse", Price: 0.0} // Price is zero, so omitted
	jsonData2, _ := json.Marshal(p2)
	fmt.Println(string(jsonData2)) 
	// Output: {"product_id":2,"productName":"Mouse"}
}
```

### Empty Structs

- `struct{}` has zero size.
- Used for signalling, not storing data.

```go
done := make(chan struct{}) // syntax for channel declaration
go func() {
    // work
    done <- struct{}{} // send signal
}()
<-done // wait
```

