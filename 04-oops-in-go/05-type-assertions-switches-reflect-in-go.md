- In Go, **interfaces** are the foundation of polymorphism. An interface variable holds:
	- **Dynamic Type** → actual concrete type assigned at runtime
	- **Dynamic Value** → the data value of that concrete type

- To _inspect_ or _extract_ this dynamic information, Go provides:  
	- **Type Assertions**  
	- **Type Switches**  
	- **Reflection** (`reflect` package)

### Type Assertions

- A **type assertion** extracts the underlying concrete value from an interface.

```go
value := iface.(ConcreteType)
```

- If the underlying type of `iface` **matches** `ConcreteType`, it succeeds. Otherwise, Go panics.

```go
var i interface{} = "hello"
s := i.(string) // succeeds
fmt.Println(s)  // hello

// Panic
var i interface{} = "hello"
n := i.(int) // panic: interface conversion: string is not int
```

- Use **comma ok idiom** to prevent panic.

```go
var i interface{} = "hello"

s, ok := i.(string)
if ok {
    fmt.Println("string:", s)
} else {
    fmt.Println("not a string")
}

n, ok := i.(int)
if !ok {
    fmt.Println("not an int") // handled safely
}
```

### Type Switches

- A **type switch** is a more elegant way to handle multiple possible underlying types.
- `i.(type)` works only with interface.

```go
func Describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case bool:
        fmt.Println("Boolean:", v)
    default:
        fmt.Printf("Unknown type %T\n", v)
    }
}

Describe(42)
Describe("Go")
Describe(true)
Describe(3.14)

// Output
// Integer: 42
// String: Go
// Boolean: true
// Unknown type float64
```

- The syntax `switch v := i.(type)` extracts both the value and its type.
- Unlike normal switches, you **can’t use type switches outside a function body** — only inside.
- `v` takes on the **concrete type** of each case branch.

### Use Case: Interfaces and Type Switch

```go
type Describer interface {
    Describe() string
}

type Person struct {
    Name string
}

func (p Person) Describe() string {
    return "Person: " + p.Name
}

type Animal struct {
    Species string
}

func (a Animal) Describe() string {
    return "Animal: " + a.Species
}

func Identify(d Describer) {
    switch v := d.(type) {
    case Person:
        fmt.Println("It's a person:", v.Name)
    case Animal:
        fmt.Println("It's an animal:", v.Species)
    default:
        fmt.Println("Unknown describer")
    }
}

func main() {
    Identify(Person{"Alice"})
    Identify(Animal{"Dog"})
}
```

### Dynamic Typing in Go

- Go is **statically typed**, but interfaces enable dynamic typing behaviour.
- At runtime, an `interface variable can point to any concrete type `— so the “dynamic type” can vary even though the interface variable is static.

```go
var any interface{}
any = 42
fmt.Printf("%T\n", any) // int

any = "Go"
fmt.Printf("%T\n", any) // string
```

- Go keeps track of this dynamic type internally.

### Reflection (Using `reflect` Package)

Reflection allows you to **inspect**, **analyse**, and **manipulate** variables at runtime.
```go
import "reflect"

// ... //

var x float64 = 3.4
fmt.Println("type:", reflect.TypeOf(x))  // float64
fmt.Println("value:", reflect.ValueOf(x)) // 3.4
```

-  `reflect.TypeOf()` 
	- Returns the **type metadata** (like `int`, `string`, `*MyStruct`).
- `reflect.ValueOf()` 
	- Returns a **reflect.Value**. we can - Inspect the actual runtime value, Modify values (if addressable)

#### Use Case - Inspect Struct Fields

```go
type User struct {
    Name string
    Age  int
}

func Inspect(i interface{}) {
    v := reflect.ValueOf(i)
    t := reflect.TypeOf(i)

    for i := 0; i < t.NumField(); i++ {
        fmt.Printf("%s = %v\n", t.Field(i).Name, v.Field(i))
    }
}

func main() {
    u := User{"Alice", 25}
    Inspect(u)
}

// output - 
Name = Alice
Age = 25
```

#### Use Case - Modifying a Value with Reflection

```go
func main() {
    x := 10
    v := reflect.ValueOf(&x).Elem() // addressable value
    if v.CanSet() {
        v.SetInt(20)
    }
    fmt.Println(x) // 20
}
```

#### Important Reflection Functions

|Function|Purpose|
|---|---|
|`reflect.TypeOf(x)`|Returns type info|
|`reflect.ValueOf(x)`|Returns runtime value|
|`v.Kind()`|Gets underlying kind (`int`, `struct`, etc.)|
|`v.Interface()`|Converts `reflect.Value` back to `interface{}`|
|`v.CanSet()`|Checks if a value is addressable/modifiable|
|`t.NumField()`|Counts struct fields|
|`t.Field(i)`|Accesses struct field metadata|

#### Reflections vs Type Assertions

| Use Case                    | Type Assertion | Reflection     |
| --------------------------- | -------------- | -------------- |
| Known possible types        | Best choice    | Overkill       |
| Unknown / arbitrary structs | Limited        | Use reflection |
| Performance                 | Fast           |  Slower        |
| Type safety                 | Compile-time   | Runtime only   |
#### Combining Reflection with Interfaces

- Reflection is often used with empty interfaces to handle arbitrary input types:

```go
func PrintDetails(i interface{}) {
    v := reflect.ValueOf(i)
    t := reflect.TypeOf(i)
    fmt.Printf("Type: %v, Kind: %v, Value: %v\n", t, t.Kind(), v)
}

PrintDetails(42)
PrintDetails("Go")
PrintDetails([]int{1, 2, 3})
```

#### Real-World Example — JSON Marshalling

The `encoding/json` package uses reflection internally to:
- Inspect struct fields
- Read tags (`json:"name"`)
- Serialize to JSON dynamically

```go
data, _ := json.Marshal(user)
```

- Go’s reflection system is analyzing your struct fields at runtime.

### Use Case: Combining all the Topics

```go
func Process(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Twice:", v*2)
    case string:
        fmt.Println("Upper:", strings.ToUpper(v))
    default:
        // fallback using reflection
        val := reflect.ValueOf(i)
        fmt.Printf("Unknown type %T with value %v\n", i, val)
    }
}

Process(10)
Process("hello")
Process([]int{1, 2, 3})

// output
// Twice: 20
// Upper: HELLO
// Unknown type []int with value [1 2 3]
```
