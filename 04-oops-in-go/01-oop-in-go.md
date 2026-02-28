- Go’s OOP is **composition-based**, not inheritance-based.
### Objects = Structs

- Structs are **the building blocks of OOP in Go**.    
- They hold fields (state).
- Struct literals can use **named fields** or **positional fields**.

```go
type Person struct {
    Name string
    Age  int
}
```

- Struct literals initialize objects

```go
p := Person{Name: "Alice", Age: 30}
```

### Methods = Receiver Functions

- Methods define behaviour for structs.    
- They are declared with as a **receiver** in GO.

```go
func (p Person) Greet() string {
    return "Hello, " + p.Name
}
```

**Pointer receiver** vs **value receiver**:
- Pointer receivers (`*Person`) allow modifying the struct or the original data-holder.    
- Value receivers (`Person`) work on a copy of the original value.

```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

### Polymorphism = Interfaces

- Define a **set of methods** for the similar objects.    
- Any type implementing those methods **satisfies the interface** automatically.
- Interfaces enable **polymorphism** without inheritance.

```go
type Greeter interface {
    Greet() string
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

p := Person{Name: "Alice"}
SayHello(p)  // Person satisfies Greeter
```

- Multiple structs can implement the same interface.

```go
type Shape interface {
    Area() float64
}

// struct
type Circle struct{ Radius float64 }
// receiver function
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

// struct
type Rectangle struct{ Width, Height float64 }
// same receiver function - satifies the interface
func (r Rectangle) Area() float64 { return r.Width * r.Height }

// common function for different structs - via interface
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

PrintArea(Circle{5})
PrintArea(Rectangle{2, 3})
```

### Inheritance / Embedding = Composition

- Go favours **composition over inheritance**.    
- Embed one struct into another to **reuse fields & methods**.

```go
type Address struct {
    City, Country string
}

type Employee struct {
    Person  // embedding
    Address // embedding
    ID      int
}

e := Employee{
    Person:  Person{Name: "Bob", Age: 25},
    Address: Address{City: "NY", Country: "USA"},
    ID:      101,
}

fmt.Println(e.Name, e.City) // accessible directly
```

### Encapsulation

- Controlled via **capitalization**:
    - **Exported (public)** → Capitalized fields/methods (`Name`, `Greet()`)        
    - **Unexported (private)** → lowercase (`age`, `calculateSalary()`)

```go
type Person struct {
    name string // private
    Age  int    // public
}
```

### Abstract Types

- Interfaces can act as **abstract types**.    
- You can declare a variable of an interface type and assign any struct that satisfies it.

```go
var g Greeter
g = Person{Name: "Alice"}
fmt.Println(g.Greet())
```

### Go’s OOP Concepts Mapped

| OOP Concept    | Go Equivalent                        |
| -------------- | ------------------------------------ |
| Class          | Struct                               |
| Object         | Struct instance                      |
| Method         | Receiver function                    |
| Inheritance    | Composition / Embedding              |
| Encapsulation  | Capitalization (exported/unexported) |
| Polymorphism   | Interfaces                           |
| Abstract Class | Interface                            |
### Constructors

- Go doesn’t have constructors. Use **factory functions** for creating objects.

```go
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}

p := NewPerson("Alice", 30)
```

### Struct Tags = Metadata

- Often used in serialization (JSON, XML, DB).

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

u := User{1, "Henry"}
jsonData, _ := json.Marshal(u)
fmt.Println(string(jsonData)) // {"id":1,"name":"Henry"}
```

