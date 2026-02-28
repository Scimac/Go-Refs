###  Pointers in Go

- A **pointer** holds the **memory address** of a variable.
- Declared with `*T` (pointer to type `T`).
- Zero value of a pointer is `nil`.

```go
var x int = 42
var p *int = &x   // p points to x

// type - int ; type of the pointer is *int

fmt.Println(x)   // 42
fmt.Println(&x)  // address of x
fmt.Println(p)   // same address
```

- Use `*p` to access the value a pointer points to - This is called **Dereferencing a Pointer**

#### Benefits of pointers

- Avoid copying large structs or arrays. 
- Allow functions to modify variables outside their scope. 
- Enable efficient memory usage.

- Avoid pointers when:
	- Using small/basic types (int, float, string).
	- Unnecessary complexity — Go is designed for simplicity.
- Avoid `unsafe.Pointer` unless doing low-level system work. ( Defined in the **`unsafe` package**. Lets you **convert between arbitrary pointer types** )
#### Passing by Value vs Passing by Reference

- Go passes arguments **by value** (copies).  
- To modify the original, pass a pointer i.r a reference to the value.

```go
func updateValue(v int) {
    v = 100
    return v
}

func updatePointer(p *int) {
    *p = 100
}

func main() {
    x := 42
    
    // Passing by value.
    v := updateValue(x)
    fmt.Println(x)  // 42 (unchanged)
    fmt.Println(v)  // 100
	
	// Passing by reference - Variable referenced gets updated
    updatePointer(&x)
    fmt.Println(x)  // 100 (modified via pointer)
}
```

#### Pointers with Structs 
#important

- Go allows shorthand: if you have a pointer to a struct, you can access fields without explicit dereferencing (`p.Age` instead of `(*p).Age`).
- Another shorthand in receiver function is - no need to pass the pointer as `&variable`, when calling a receiver function.

```go
type Person struct {
    Name string
    Age  int
}

// Shorthand 1
func birthday(p *Person) {
    p.Age++ // go shorthand, same as (*p).Age++
}

// shorthand 2
func (p *Person) birthday () {
	p.Age++
}

func main() {
    john := Person{"John", 30}
    birthday(&john) // normal use
    
    fmt.Println(john.Age) // 31
    
    john.birthday() // no need to use & operator
    fmt.Println(john.Age) // 32
}
```

### `new` vs `make`

- `new(T)` → allocates memory, returns `*T` (`zero-initialized`).
- `make` → used only for **slices, maps, channels** (returns `initialized value`, not a pointer).

```go
// Initialising an empty pointer
p := new(int)    // *int, zero value (0)
fmt.Println(*p)  // 0

*p = 10
fmt.Println(*p)  // 10

// creating a data structure using value returned by make()
m := make(map[string]int)
m["a"] = 1
```