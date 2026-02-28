## Error Handling in Go

- In Go, errors are values, not exceptions. The built-in `error` type is an **interface**:
```go
type error interface {
    Error() string
}
```
- Functions that may fail usually return `(value, error)`.

### Example

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

### Creating Errors

```go
// Using errors.New
err := errors.New("something went wrong") // import "errors"
```

```go
// Using fmt.Errorf
err := fmt.Errorf("file %s not found", "data.txt")
```

### Checking Errors

- Always check `if err != nil` before using the result.
```go
val, err := divide(4, 2)
if err != nil {
    // handle error
}
```

### Custom Errors

```go
// define your own error types for more context:
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}

func riskyOperation() error {
    return &MyError{Code: 404, Msg: "Resource not found"}
}

func main() {
    err := riskyOperation()
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### Sentinel Errors (Predefined)

```go
// Declare common reusable errors.
var ErrNotFound = errors.New("not found")

func lookup(key string) (string, error) {
    if key != "foo" {
        return "", ErrNotFound
    }
    return "bar", nil
}

func main() {
    _, err := lookup("baz")
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Key not found")
    }
}
```

### Error Wrapping

- Use `fmt.Errorf` with `%w` to wrap errors.
- Use `errors.Is` or `errors.As` to inspect.

```go
var ErrPermission = errors.New("permission denied") // import "errors"

func openFile() error {
    return fmt.Errorf("open failed: %w", ErrPermission)
}

func main() {
    err := openFile()
    if errors.Is(err, ErrPermission) {
        fmt.Println("Permission issue detected")
    }
}
```

- Wrap errors with context (`%w`) for debugging.

### Panic and Recover
#### Panic

- Used for unrecoverable errors (e.g., out of bounds, nil dereference).
- Stops normal execution.

```go
func main() {
    panic("something went very wrong")
}
```
#### Recover

- Used to regain control after a panic.
- Must be called inside `defer`.

```go
func safe() {
    defer func() {
        if r := recover(); r != nil { // shorthand for declaration and reuse
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("danger!")
}

func main() {
    safe()
    fmt.Println("Program continues...")
}
```
- Only use `panic` for programmer errors (not user input issues).
- Return `error` values instead of panicking, unless it’s truly unrecoverable.


