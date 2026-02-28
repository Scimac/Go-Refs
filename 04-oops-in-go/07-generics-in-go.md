- Before Go 1.18, developers often used the **empty interface (`interface{}`)** to handle different types
- This sacrificed **type safety** and required **manual type assertions**.

- Go Generics solve that by allowing you to write **functions, methods, and types** that work with any type — while still being **type-safe and compiled**.

### Generics - Syntax

- The generic type parameter list is defined using square brackets `[]`.

```go
func FunctionName[T TypeConstraint](args...) ReturnType { ... }
```

-  `T` = **type parameter name** (like a variable, but for types)
- `TypeConstraint` = allowed types or an interface defining allowed operations

```go
// Non-generic function
func SumInt(a, b int) int {
    return a + b
}

// Generic function
func Sum[T int | float64](a, b T) T {
    return a + b
}

fmt.Println(Sum(10, 20))     // 30
fmt.Println(Sum(2.5, 3.7))   // 6.2
```

- `T` is a **type parameter**.  
- `int | float64` is a **type constraint** — it restricts what types `T` can be.

#### Use Case: Generic Print Function

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"Go", "Rust", "C++"})
```

> `any` means “any type” — equivalent to `interface{}` but type-safe.

### Type Constraints - Declarations

- Constraints define what operations are allowed on type parameters.

#### Specific Allowed Types

```go
// You can only pass `int` or `float64`
func AddNumbers[T int | float64](a, b T) T {
    return a + b
}
```

#### Using `~` for Type Approximation

```go
// The `~` operator allows **underlying type matching**.
type MyInt int

func Double[T ~int](v T) T {
    return v * 2
}

var x MyInt = 5
fmt.Println(Double(x)) // 10 - works because MyInt’s underlying type is int
```

#### Constraint Interfaces

```go
// Constraints can also be defined as **interfaces** that specify methods or operations.

type Adder[T any] interface {
    Add(a, b T) T
}
```

#### Built-in `comparable` Constraint

- Go provides a built-in constraint called **`comparable`**, for types that support `==` and `!=`.
```go
func IndexOf[T comparable](slice []T, item T) int {
    for i, v := range slice {
        if v == item {
            return i
        }
    }
    return -1
}

fmt.Println(IndexOf([]int{1, 2, 3}, 2))       // 1
fmt.Println(IndexOf([]string{"a", "b"}, "b")) // 1
```

#### Type Inference

- In many cases, Go **infers type parameters** automatically.

```go
func Identity[T any](x T) T { return x }

fmt.Println(Identity(42))     // infers T as int
fmt.Println(Identity("Go"))   // infers T as string
```
### Generic Structs

- You can define structs that are parameterized by type.

```go
type Pair[T any, U any] struct {
    First  T
    Second U
}

p := Pair[int, string]{First: 10, Second: "Go"}
fmt.Println(p) // {10 Go}
```

### Generic Methods
- Structs with type parameters can also have methods using those parameters.

```go
type Container[T any] struct {
    items []T
}

func (c *Container[T]) Add(item T) {
    c.items = append(c.items, item)
}

func (c Container[T]) GetAll() []T {
    return c.items
}

func main() {
    var ints Container[int]
    ints.Add(10)
    ints.Add(20)
    fmt.Println(ints.GetAll()) // [10 20]
}
```

### Generic Type Aliases

```go
type Number interface {
    int | int64 | float64
}

func SumAll[T Number](nums ...T) T {
    var total T
    for _, v := range nums {
        total += v
    }
    return total
}
```

### Nested Generics

- Generics can be nested in structs and methods.

```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

func main() {
    n1 := Node[int]{Value: 10}
    n2 := Node[int]{Value: 20, Next: &n1}
    fmt.Println(n2.Next.Value) // 10
}
```

### Combining Interfaces with Generics

 - You can use generic type constraints in combination with interfaces.
```go
type HasArea[T any] interface {
    Area() T
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func PrintArea[T HasArea[float64]](s T) {
    fmt.Println("Area:", s.Area())
}

PrintArea(Circle{Radius: 5})
```

### Real-World Example

#### Generic Map Function

```go
func Map[T any, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

nums := []int{1, 2, 3, 4}
doubled := Map(nums, func(x int) int { return x * 2 })
fmt.Println(doubled) // [2 4 6 8]
```

#### Generic Stack

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    n := len(s.items)
    item := s.items[n-1]
    s.items = s.items[:n-1]
    return item
}

func (s Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func main() {
    var stack Stack[string]
    stack.Push("Go")
    stack.Push("Lang")
    fmt.Println(stack.Pop()) // Lang
}
```

#### Generics with Constraints Packages

- Go’s `constraints` package (`golang.org/x/exp/constraints`) provides useful predefined constraints.

```go
import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Works for all **ordered types** — integers, floats, and strings.
```

### Limitations of Go Generics

-  Some operations are **not supported**:
	- You **cannot** perform arithmetic on arbitrary types (must be constrained).
	- **No operator overloading.**
	- **No runtime type reflection** for type parameters.
	- **Type parameters cannot be used in composite literal keys** (e.g., `map[T]U` with unconstrained `T`).
