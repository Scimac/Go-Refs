## Memory Layout & Escape Analysis in Go (Deep Dive)

### Go Memory Model — Big Picture

- Go manages memory automatically using:
	- **Stack** 
	- **Heap**
	- **Garbage Collector (GC)**

- You **do not choose** where variables live.  
- The **compiler decides**, using **escape analysis**.

## Stack vs Heap (Conceptual)

### Stack Memory

- Fast allocation    
- Function-scoped
- Automatically freed when function returns
- No GC involvement
- Limited size

### Heap Memory

- Slower allocation
- Shared across goroutines
- Managed by GC
- Used for long-lived data
- Larger but GC-heavy

### Key Rule
#InterviewQs #important 

> **If a value can be proven to not outlive its function, it stays on the stack.  
> If it escapes the function, it moves to the heap.**

This decision is made at **compile time**, not runtime.

## Escape Analysis — What Is It?

- **Escape analysis** is a compiler optimization that determines:

> “Does this variable need to live beyond this function?”
> If yes → **heap**  , If no → **stack**

### Simple Example — Stack Allocation

```go
func add(a, b int) int {
    c := a + b
    return c
}
```

- `a`, `b`, `c` live on the **stack**  
- No pointers escaping  
- No GC involvement
### Pointer Escape (MOST IMPORTANT)
#important #InterviewQs 

```go
func makePointer() *int {
    x := 10
    return &x
}
```

- `x` **escapes** the function.

- Why?
	- Returning `&x`
	- Caller needs it after function returns

> `x` is allocated on the **heap**, not stack.

### Heap Allocation WITHOUT Pointers

```go
func f() {
    s := make([]int, 1000)
}
```

- This slice **may escape** depending on usage.

- Why?
	- Slices internally contain a pointer
	- If the slice header escapes, underlying array escapes
	
> Compiler decides.

### Slice Escape Analysis (Critical for DSA)
#InterviewQs 

#### Case 1: Slice stays on stack

```go
func sum() int {
    s := []int{1, 2, 3}
    return s[0] + s[1]
}
```

- Slice does NOT escape  
- Array allocated on stack

#### Case 2: Slice escapes

```go
func makeSlice() []int {
    s := []int{1, 2, 3}
    return s
}
```

- Slice escapes → underlying array on heap

### Maps ALWAYS Live on Heap
#important 

```go
func f() {
    m := make(map[string]int)
    m["x"] = 1
}
```

- Maps are **always heap allocated**, Because:
	- Runtime-managed hash table
	- Internal pointers
	- Needs GC tracking

### Struct Allocation: Stack or Heap?
#InterviewQs 

```go
type User struct {
    Name string
    Age  int
}

func f() {
    u := User{Name: "A", Age: 10} // Stack allocated
}

func newUser() *User {
    u := User{Name: "A", Age: 10}
    return &u // escapes → heap
}
```

### Interface Escapes (Subtle & Important)
#important 

```go
func log(v interface{}) {
    fmt.Println(v)
}

func f() {
    x := 10
    log(x)
}
```

- `x` may escape because:
	- Interface holds `(type, value)`
	- Value might be boxed
- Interfaces are a **common cause of heap allocations**

### Closures & Escape
#important 

```go
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}
```

- `x` escapes
	- Used after function returns
	- Closure captures it
-  Heap allocation

## Escape Analysis Tooling (VERY IMPORTANT)
#important #InterviewQs 

- How to See Escapes

```shell
go build -gcflags="-m"
```

- **Example:**

```go
func f() *int {
    x := 10
    return &x
}

// output: x escapes to heap
```

## Garbage Collection (GC) Interaction

- **What GC Tracks**
	- Heap objects only
	- Reachable objects
	- Pointers between objects
- **GC Does NOT Track**
	- Stack values
	- Compile-time known lifetimes

## Why Escape Analysis Matters

- **Performance**
	- Heap allocation = slower
	- GC pressure increases
	- Cache locality worse

```go
func loop() {
    for i := 0; i < 1_000_000; i++ {
        x := new(int)
        *x = i
    }
}

// Massive heap churn
```

## How to Reduce Escapes
#important 

1. Prefer Values Over Pointers
```go
func f(u User) {} // better
func f(u *User) {} // only if mutation needed
```

2. Avoid Returning Pointers to Locals
```go
func newUser() User {
    return User{Name: "A"} // Allows stack allocation
}
```

3. Minimize Interface Usage in Hot Paths
```go
func process(i int) {}        // faster
func process(i interface{}) {} // slower
```

4. Reuse Memory (sync.Pool)
```go
var pool = sync.Pool{
    New: func() any {
        return make([]byte, 1024)
    },
}
```

## Summary

```txt
STACK
 ├─ local variables
 ├─ function arguments
 └─ non-escaping structs, arrays

HEAP
 ├─ escaping variables
 ├─ returned pointers
 ├─ maps
 ├─ slices that escape
 ├─ closures
 └─ interface-boxed values
```

