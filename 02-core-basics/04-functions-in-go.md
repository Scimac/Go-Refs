### Basic Functions Declaration

```go
func add(a int, b int) int { // type for both aregument and return needed
    return a + b
}
```

- if there are multiple arguments with same type - 
```go
func multiply(a, b int) int { // type can be clubbed for the arguments
    return a * b
}
```
### Function with Multiple return values 

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divmod(10, 3)
fmt.Println(q, r) 
// 3 1
```

- multiple return variables can be named - 
```go
func rectangle(w, h int) (area int, perimeter int) {
    area = w * h
    perimeter = 2 * (w + h)
    return // not really recommended, can cause confusion in large functions 
    // return area, perimeter
}

// Use **multi-return** instead of throwing errors.
```
- Named returns act like variables inside the function.
### Variadic Functions

- Variadic function are type of functions that can work with varied number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4)) // 10
```

- - `...int` means any number of `int` arguments. (Just like rest operator in JS)
- Internally received as a **slice**.
### Function Literals (Anonymous Functions/Lambdas)

- Function literals are nameless functions.

```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4)) // 7
```

- These functions can be triggered immediately. (like - IIFE  in JS)

```go
func(li string) string {
	time.Sleep(5 * time.Second)
	return li
}(l) // l is the actual value passed
```
### Higher Order Functions

```go
func operate(a, b int, f func(int, int) int) int {
    return f(a, b)
}

// Examples of anonymous functions
add := func(x, y int) int { return x + y }
mul := func(x, y int) int { return x * y }

fmt.Println(operate(3, 4, add)) // 7
fmt.Println(operate(3, 4, mul)) // 12

// since function type declaration can get bigger and error-prone
// use Type aliases
type fnType func(int, int) int

func coOperate(a, b int, f fnType) int {
    return f(a, b)
}
```

### Methods (Functions with Receivers)

- A **method** is a function attached to a type.

```go
type Rectangle struct {
    W, H int
}

// Value receiver 
// (Passed by value. It is a Copy. Does not affect orginal)
func (r Rectangle) Area() int {
    return r.W * r.H
}

// Pointer receiver 
// (Address/Pointer to original value is passed. Modifies original)
func (r *Rectangle) Scale(factor int) {
    r.W *= factor
    r.H *= factor
}
// Use pointer receivers when modifying a struct.

func main () {
	rect := Rectangle{2, 3}
	fmt.Println(rect.Area()) // 6
	rect.Scale(2)
	fmt.Println(rect.Area()) // 24
}
```
### Defer in Functions

- `defer` delays execution until function exits.
- Useful for closing files, unlocking `mutexes`, etc.

```go
func demo() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
demo()
// hello
// world
```
### Function Pointers

- Functions are values -- you can assign and call dynamically.

```go
add := func(x, y int) int { return x + y }

var op func(int, int) int
op = add
fmt.Println(op(2, 3)) // 5
```
### Function Recursion

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
fmt.Println(factorial(5)) // 120
```

### Error Handling with Functions

- Go doesn’t use exceptions.
- Functions often return `(value, error)`.

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
```