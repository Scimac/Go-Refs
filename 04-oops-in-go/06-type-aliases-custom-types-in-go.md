- Go allows you to **define new types** based on existing ones using the `type` keyword.

```go
type Age int
type Name string
```

- These are **new, distinct types**, even though they’re based on built-in ones.  
- They don’t automatically inherit the methods of their underlying type.

```go
var a Age = 30
var n Name = "Alice"

// var x int = a // cannot use a (Age) as int

// You must explicitly convert between the custom type and the base type:
var x int = int(a)
```

### Why Create Custom Types?

Custom types are used to:  
- Improve **readability** and express intent  
- Add **type safety** (e.g., prevent mixing unrelated values)  
- Attach **methods** to non-struct types  
- Represent **domain-specific concepts**

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// ...
func main () {
	c := Celsius(30)
	fmt.Println(c.ToFahrenheit()) // 86°F
}
// ...
```

### Type Definition vs Type Alias

| Keyword             | Example            | Behavior                                     |
| ------------------- | ------------------ | -------------------------------------------- |
| **Type Definition** | `type MyInt int`   | Creates a _new distinct type_                |
| **Type Alias**      | `type MyInt = int` | Just a _different name_ for an existing type |
```go
type MyInt int   // new type
type YourInt = int // alias

var a MyInt = 10
var b YourInt = 10

// var c int = a // needs conversion
var d int = b    // alias acts as int
```

### Defining Methods on Custom Types

- Any **named type** (non-interface) can have methods — not just structs!
- The receiver type can be a **value** or **pointer** — just like struct methods.

```go
type Counter int

func (c *Counter) Increment() {
    *c++
}

func (c Counter) Value() int {
    return int(c)
}

func main() {
    var c Counter = 10
    c.Increment()
    fmt.Println(c.Value()) // 11
}
```

- Methods can only be defined on **named types** — not aliases.

```go
type MyInt int // OK
func (m MyInt) Double() int { return int(m) * 2 }

// type YourInt = int // cannot define methods on alias
```

- You can wrap any built-in type to add functionality.

```go
type StringList []string

func (s StringList) Join() string {
    return strings.Join(s, ", ")
}

func main() {
    list := StringList{"Go", "Rust", "C++"}
    fmt.Println(list.Join()) // Go, Rust, C++
}
```

### Methods and Interfaces

- Custom types can implement interfaces, even if they are based on primitive types.

```go
type MyString string

func (s MyString) Len() int {
    return len(s)
}

type Sizer interface {
    Len() int
}

func PrintSize(s Sizer) {
    fmt.Println("Length:", s.Len())
}

func main() {
    var str MyString = "Hello"
    PrintSize(str) // Length: 5
}
```


- Same rules as for structs apply for type of receivers:

| Receiver             | Description           | When to Use                                                   |
| -------------------- | --------------------- | ------------------------------------------------------------- |
| **Value Receiver**   | Works on a copy       | When the method doesn’t modify the value                      |
| **Pointer Receiver** | Works on the original | When the method modifies or needs to avoid copying large data |
### Zero Values and Initialization

- Custom types follow Go’s **zero-value rule**.  
- Each base type’s zero value applies to the custom type as well

```go
type MyInt int
type MySlice []string
type MyMap map[string]int

var i MyInt
var s MySlice
var m MyMap

fmt.Println(i) // 0
fmt.Println(s) // nil
fmt.Println(m) // nil
```

### Embedding Custom Types in Structs

- Custom types can be embedded inside structs to extend functionality.

```go
type Distance float64
func (d Distance) Miles() float64 { return float64(d) * 0.621371 }

type Journey struct {
    Distance
}

func main() {
    j := Journey{Distance: 100}
    fmt.Println(j.Miles()) // 62.1371
}
```

### Example: Domain-Oriented Design

- Here, `Money` adds **domain meaning**, and methods enforce **type safety** and **clarity**.

```go
type Money float64

func (m Money) String() string {
    return fmt.Sprintf("$%.2f", m)
}

type Wallet struct {
    balance Money
}

func (w *Wallet) Deposit(amount Money) {
    w.balance += amount
}

func (w Wallet) Balance() Money {
    return w.balance
}

func main() {
    w := Wallet{}
    w.Deposit(50)
    fmt.Println("Balance:", w.Balance()) // $50.00
}
```

