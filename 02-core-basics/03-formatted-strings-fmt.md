- The `fmt` package helps in working with formatted strings.
-  It implements formatted I/O (input/output). It’s used for printing to console, creating formatted strings, and scanning input.

## Printing
### Basic Printing

```go
import "fmt"

func main() {
    fmt.Print("Hello ")         // no newline
    fmt.Println("World")        // adds newline
    fmt.Printf("Value: %d\n", 5) // formatted print
}
```
### Creating Strings

- We can format strings for assigning it to variables.
```go
msg := fmt.Sprintf("Pi = %.2f", 3.14159)
fmt.Println(msg) // "Pi = 3.14"
```
- `Sprint`, `Sprintln` → string equivalents of `Print`, `Println`.

### Common Formatting Verbs
|Verb|Description|Example|
|---|---|---|
|`%v`|Default format|`fmt.Printf("%v", 42)` → `42`|
|`%+v`|Struct with field names|`fmt.Printf("%+v", person)`|
|`%#v`|Go-syntax representation|useful for debugging|
|`%T`|Type of value|`"int"`, `"string"`|
|`%%`|Literal percent sign|`"50%%"`|
### Numeric Formatting
|Verb|Description|Example|
|---|---|---|
|`%d`|Decimal integer|`255`|
|`%b`|Binary|`11111111`|
|`%o`|Octal|`377`|
|`%x` / `%X`|Hex (lower/upper)|`ff` / `FF`|
|`%f`|Float (decimal)|`3.141593`|
|`%.2f`|Float with precision|`3.14`|
|`%e` / `%E`|Scientific notation|`1.23e+03`|
|`%g`|Compact float (auto chooses `%f`/`%e`)|`1234.5`|
```go
num := 255
fmt.Printf("%d %b %o %x\n", num, num, num, num)
// 255 11111111 377 ff
```
### String & Character Formatting
|Verb|Description|Example|
|---|---|---|
|`%s`|String|`"Hello"`|
|`%q`|Double-quoted string|`"Hello"`|
|`%x`|Hex dump of string bytes|`68656c6c6f`|
|`%c`|Character (Unicode)|`'A'`|
```go
str := "Go"
fmt.Printf("%s %q %x\n", str, str, str)
// Go "Go" 476f
```

### Boolean & Pointers
|Verb|Description|Example|
|---|---|---|
|`%t`|Boolean|`true`, `false`|
|`%p`|Pointer address|`0xc0000140a8`|
```go
flag := true
fmt.Printf("%t\n", flag) // true

ptr := &flag
fmt.Printf("%p\n", ptr)  // 0xc...
```

### Width, Alignment, and Padding

```go
fmt.Printf("%6d\n", 42)  // "    42"
fmt.Printf("%-6d\n", 42) // "42    "
fmt.Printf("%06d\n", 42) // "000042"
```

### Structs & Slices

```go
type Person struct {
    Name string
    Age  int
}

p := Person{"Alice", 30}
fmt.Printf("%v\n", p)   // {Alice 30}
fmt.Printf("%+v\n", p)  // {Name:Alice Age:30}
fmt.Printf("%#v\n", p)  // main.Person{Name:"Alice", Age:30}
```

## Scanning
### Input with `fmt`

```go
var name string
var age int
fmt.Scan(&name, &age)  // input: "Bob 25"
fmt.Println(name, age) // Bob 25
```
- We need to pass pointers in the `Scan` function.

```go
var x int
fmt.Sscanf("age=42", "age=%d", &x) 
// Scans with format and stores values in variable
fmt.Println(x) // 42
```

## Reference

Refer [this](https://pkg.go.dev/fmt@go1.25.1) for more information.