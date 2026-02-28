### Variable Declaration & Initialization

#### Explicit Declaration

- `var` keyword is used. Type is Explicitly mentioned.
```go
var x int // explicit declaration
x = 10 // initialization

// OR can be done in one go
var y int = 20

// constant declaration
const Pi = 3.14159
const Greeting = "Hello, Go!"

// typed vs untyped
const X = 10      // untyped → flexible
var y float64 = X // X adapts type

// `iota` (Constant Generator) - Auto-incrementing counter for enums.
const (
    Red = iota
    Green
    Blue
)
fmt.Println(Red, Green, Blue) // 0 1 2

// Constants can be defined with arithmetic or bitwise operations.
// just value of the const should be computed before hand
const (
    KB = 1 << (10 * iota) // 1 KB = 1024
    MB                    // 1 MB = 1048576
    GB                    // 1 GB
)
```
#### Implicit Declaration

- Type is inferred by the GO compiler based on value assigned. Not defined by user.
```go
var test = "checked"

a := "GoLang"  // shorthand, inside functions only
b := 3.14      // float64
c := true      // boolean
```

> Cannot use `:=` outside of a function (e.g., at package level).

#### Multiple Declarations

```go
var x, y, z int = 1, 2, 3
a, b := "hello", "world"
```
#### `new` vs `make` (Memory Allocation)

- **`new`**: allocates memory for a value, returns a pointer.
```go
p := new(int)   // *int, initialized to 0
*p = 10
```

- **`make`**: used for slices, maps, and channels only.
```go
s := make([]int, 5)     // slice of length 5
m := make(map[string]int)
c := make(chan int)
```
- `new` gives pointer to zero value; `make` gives usable initialised data structure.
#### Null Values

- If the variable is only declared and not initialised, go compiler assigns `null` values to the variable.

| Type                              | Zero Value |
| --------------------------------- | ---------- |
| int, float                        | `0`, `0.0` |
| string                            | `""`       |
| bool                              | `false`    |
| pointer                           | `nil`      |
| slice, map, chan, func, interface | `nil`      |
- `nil` is only valid for **reference types** (pointer, slice, map, chan, func, interface).
- Numeric, string, and bool types **never use nil**, only zero values.
#### Shadowing Variables

- A **new variable with the same name** can be declared inside a narrower scope.
- Can lead to bugs
```go
x := 5
if true {
    x := 10 // shadows outer x
    fmt.Println(x) // 10
}
fmt.Println(x) // 5
```

- Package-level variables exist for the entire program.
- Function variables exist until the function exits.
- Closures can **capture variables**:
```go
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```
#### Blank Identifier (`_`)

- Use `_` to **ignore values** you don’t need.
```go
val, _ := strconv.Atoi("123") // ignore error
fmt.Println(val)
```

#### Numeric Literals & Underscores

```go
// can improve readability with `_` inside numbers:
x := 1_000_000 // same as 1000000

// Different bases supported
bin := 0b1010    // binary (10)
oct := 0o12      // octal (10)
hex := 0xA       // hexadecimal (10)
```

### Basic Data types in GO

- **Numeric**
    - Integers: `int`, `int8`, `int16`, `int32`, `int64` - Signed Integers
    - Unsigned: `uint`, `uint8` (`byte`), `uint16`, `uint32`, `uint64`
    - Rune: An alias for `int32`, represents a Unicode code point (i.e., a single character),(e.g., 'a', 'ñ', '世')
    - Floating point: `float32`, `float64`
    - Complex: `complex64`, `complex128` - E.g. `c := complex(2, 3)  // 2+3i`

- **String**    
    - **Immutable sequences** of bytes (`string`).

- **Boolean**    
    - `true` or `false`.

#### Immutable Strings & Rune Type

- Strings are **immutable** in Go. You cannot change characters directly.
```go
s := "hello"
// s[0] = 'H'  // ❌ error
```

- Use `[]rune` or `[]byte` for manipulation:
```go
r := []rune(s)
r[0] = 'H'
s = string(r)
fmt.Println(s) // "Hello"
```

- `rune` = alias for `int32`, represents a Unicode code point.
- `byte` = alias for `uint8`, represents raw ASCII/UTF-8 bytes.
### Composite Types in GO

- **Array**    
    `var arr [3]int = [3]int{1, 2, 3}     // Length, type is fixed`

- **Slice**    
    `nums := []int{1, 2, 3, 4}     // Length is variable, only type is fixed`

- **Struct**    
    ```go
	type Person struct {
		Name string     
		Age  int 
	} // keys are always string
	
	p := Person{"Alice", 25}
	```

- **Map**    
    `m := map[string]int{"a": 1, "b": 2}     // type of key,value is fixed`

- **Pointer**    
    ```go
    var ptr *int 
    i := 42 
    ptr = &i 
    fmt.Println(*ptr) // dereference
    ```

- **Interface**    
    `var x interface{} x = "hello" x = 42`

