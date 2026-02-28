# Numeric, String & Boolean Data Types in Go

## Numeric Data Types in Go

- Go has **explicit, strongly typed numeric categories**.
- There is **no implicit type promotion** like in C/C++ or JS.

```go
var a int = 10
var b int64 = 100
var u uint = 10
```
### Signed Integers

| Type    | Size                              |
| ------- | --------------------------------- |
| `int`   | platform dependent (32 or 64 bit) |
| `int8`  | 8 bit                             |
| `int16` | 16 bit                            |
| `int32` | 32 bit                            |
| `int64` | 64 bit                            |
### Unsigned Integers

- Unsigned integers cannot represent negative values

| Type     | Size               |
| -------- | ------------------ |
| `uint`   | platform dependent |
| `uint8`  | 8 bit              |
| `uint16` | 16 bit             |
| `uint32` | 32 bit             |
| `uint64` | 64 bit             |
### Floating-Point Types
| Type      | Size             | Use Case                                        |
| --------- | ---------------- | ----------------------------------------------- |
| `float32` | 32 bit           | Only used when space is tight                   |
| `float64` | 64 bit (default) | **float64 is preferred** — faster on most CPUs. |
```go
var f float64 = 3.14
```

- Floating-point math is **approximate**, not exact. #InterviewQs 
```go
fmt.Println(0.1 + 0.2) // 0.30000000000000004
```

### Epsilon Comparison
#InterviewQs 

- Float comparison are tricky. Be careful about them.
```go
a := 0.1 + 0.2
b := 0.3

if a == b { // This will likely be false, as math approx.
	fmt.Println("Equal")
} else {
	fmt.Println("Not Equal")
}

// Output: Not Equal
```

- To fix this, define an `epsilon` value (a small tolerance, e.g., `1e-9`)
- Compare the absolute difference of the two numbers against this tolerance using the `math.Abs` function.

```go
import "math"

const epsilon = 1e-9 // Choose an appropriate epsilon for your application

func areApproximatelyEqual(a, b float64) bool {
    return math.Abs(a - b) <= epsilon
}

// ....

if areApproximatelyEqual(a, b) {
	fmt.Println("Equal (within tolerance)")
} else {
	fmt.Println("Not Equal")
}
```

### Complex Numbers (Rare but Exists)

- Mostly used in scientific computing.
```go
var c complex128 = complex(1, 2)
fmt.Println(real(c), imag(c))
```

### Numeric Type Rules
#important 

1. No Implicit Conversion
```go
var a int = 10
var b int64 = a // compile error

// Must convert explicitly:

var b int64 = int64(a)
```

2. Mixed-Type Operations Not Allowed
```go
var a int = 10
var b float64 = 2.5

c := a + b // compile error, Explicit conversion required.
```

## Byte & Rune Type in Go
#important 

```go
type byte = uint8 // raw binary / ASCII
type rune = int32 // Unicode code point
```

```go
var c byte = 'A'
var r rune = '世'
```

- Use `rune` when working with Unicode characters. #InterviewQs 
```go
s := "世"
// len(s) != 1
fmt.Println(len(s)) // 3 (bytes)
```

```go
// Correct way
runes := []rune(s)
fmt.Println(len(runes)) // 1
```

## String Data Type in Go

- Immutable sequence of **bytes**
- UTF-8 encoded
- Value type (but reference-like internally)

```go
s := "hello"
```

### String Internals

- Internally string is a HEAD to the actual value.
- Data is immutable
- Multiple strings can share memory safely
```go
type string struct {
    ptr *byte
    len int
}
```

- Sometime, single character in string can take multiple bytes.
```go
s := "世"
fmt.Println(len(s)) // 3 (bytes)
```

### Iterating Over Strings
#important #InterviewQs 

- Incorrect (byte-wise) - Some character can take more than one byte.
```go
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
```

- Correct (Unicode-safe) - `range` over string iterates over **runes**, not bytes.
```
for _, r := range s {
    fmt.Println(r)
}
```

- If want to count the characters only, convert string into slice of runes -
```go
runes := []rune(s)
fmt.Println(len(runes)) // 1
```

### String Operations

#### Concatenation

- Repeated concatenation in loops is inefficient. Use `strings.Builder` instead.
```go
// inefficient
s := "go" + "lang"

// string.Builder
var b strings.Builder
b.WriteString("go")
b.WriteString("lang")
```

#### String Immutability

```go
// Strings cannot be modified.

s := "hello"
s[0] = 'H' // ❌ compile error
```

## Boolean (`bool`) Data Type

- Only values: `true`, `false`
- Zero value: `false`

```go
var flag bool = true
```

### No Truthy / Falsy
#important 

```go
if 1 { }       // invalid
if "hello" { } // invalid
```

- Must be explicit boolean
```go
if x > 0 { }
```

- **Go does NOT allow implicit boolean evaluation.** #InterviewQs 

### Boolean Operators

```go
&&  // logical AND
||  // logical OR
!   // NOT
```

### Comparison Operators

```go
== != < > <= >=
```

```go
// String comparison
"apple" < "banana" // lexicographic

// Boolean Comparison
true == false // valid
true > false  // invalid
```

## Zero Value for All Data Types

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| float64 | 0.0        |
| string  | ""         |
| bool    | false      |
```go
var s string
var n int
var b bool

fmt.Println(s, n, b) // "", 0, false
```

