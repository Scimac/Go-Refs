### Interface - Introduction

- An **interface** defines a set of method signatures (receiver functions attached).  
- Any type that implements all those methods **implicitly satisfies** the interface — no `implements` keyword required.

```go
// if you have only one receiver function
// the name of interface is often named as - '[receiver_fn] + er'
type Describer interface {
    Describe() string
}

type Person struct {
    Name string
}

func (p Person) Describe() string {
    return "Person named " + p.Name
}

func main() {
    var d Describer
    d = Person{Name: "Alice"} // Person implements Describer
    fmt.Println(d.Describe())
}
```
### Key Characteristics of Interfaces

|Property|Description|
|---|---|
|**Implicit implementation**|No `implements` keyword|
|**Behavior-driven**|Focused on _what_ a type can do, not _what it is_|
|**Decoupling mechanism**|Allows functions to depend on behavior, not concrete types|
|**Zero value**|`nil`|
|**Dynamic type and value**|Holds a pair: _(concrete value, concrete type)_|
### Interface Method Sets

- A **method set** defines what methods a type exposes to interfaces.

```go
type Reader interface {
    Read() string
}

type FileReader struct{}

func (FileReader) Read() string { return "reading file..." }
```

- `FileReader` satisfies `Reader` because it implements `Read()`.
#### Pointer vs Value Receivers

| Method Receiver                | Type That Satisfies Interface |
| ------------------------------ | ----------------------------- |
| Value receiver (`func (T)`)    | Both `T` and `*T`             |
| Pointer receiver (`func (*T)`) | Only `*T`                     |
```go
type Counter struct{ val int }

func (c *Counter) Increment() { c.val++ }

type Incrementer interface {
    Increment()
}

var i Incrementer
c := Counter{}

// i = c   error — value type doesn't have Increment()
// i = &c  pointer has the method
```

### Empty Interface (`interface{}` / `any`)

- The **empty interface** can hold _any type_ — because all types implement zero methods.
- Since Go 1.18, `interface{}` and `any` are **identical**.

```go
func PrintAny(value interface{}) {
    fmt.Println(value)
}

PrintAny(42)
PrintAny("hello")
PrintAny([]int{1, 2, 3})

func DoSomething(x any) { ... }
```

### Interface Embedding

- Interfaces can embed other interfaces to compose behavior.

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

- Any type implementing both `Read()` and `Write()` automatically satisfies `ReadWriter`.

### Polymorphism in Go

 Interfaces enable **runtime polymorphism** — the ability to call methods on values without knowing their concrete type.

```go
type Shape interface {
    Area() float64
}

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

PrintArea(Circle{5})
PrintArea(Rectangle{2, 3})
// Both `Circle` and `Rectangle` are treated as `Shape`.
```

### Nil Interfaces

- Two scenarios of Interfaces being `nil` -
	- Case 1: Nil interface variable
	  ```go
	    var s Shape
		fmt.Println(s == nil) // true
		```
	- Case 2: Non-nil interface with nil concrete value
	  ```go
		  var c *Circle = nil
		  var s Shape = c
		  fmt.Println(s == nil) // false
		```

- Reason: `s` holds type = `*Circle`, value = `nil`.  The interface itself is non-nil.

### Interface Values Internally
#important 

- An interface value holds two things:
```md
(interface value)
---------------------
| dynamic type      |
| dynamic value     |
---------------------
```

```go
var i Shape = Circle{5}
```

- Dynamic type = `Circle`
- Dynamic value = `{5}`

### Best Practices

- Design **small interfaces** — one or two methods.  
- Rely on **behaviour**, not concrete types.  
- Avoid overusing `interface{}` (use generics when possible).  
- Prefer **composition** — combine simple interfaces instead of large ones.  
- Keep interface definitions near their use sites (don’t centralise all).

### Real-World Examples

#### `fmt.Stringer`

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    p := Person{"Alice", 30}
    
    // Because `fmt` uses the `Stringer` interface internally.
    // Automatically calls p.String()
    fmt.Println(p) // here
}

// output: Alice (30 years old)
```

#### Composing Interfaces

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}

type File struct{}

func (File) Read() string      { return "Reading file" }
func (File) Write(s string)    { fmt.Println("Writing:", s) }

func Process(rw ReadWriter) {
    fmt.Println(rw.Read())
    rw.Write("Data")
}

func main() {
    f := File{}
    Process(f)
}
```

