### Memory Management in Go

#### Key Features

- **Automatic allocation**
    - Variables are allocated on the **stack** (local, short-lived) or **heap** (long-lived).
    - Go decides where, based on escape analysis.
	
- **Escape Analysis**
    - If a variable "escapes" the function (e.g., returned by pointer), Go allocates it on the **heap**.
		```go
		func makePointer() *int {
		    x := 42   // normally stack
		    return &x // escapes, so heap
		}
		```
	
- **Garbage Collection (GC)**
    - Go has a `concurrent, tri-color mark-and-sweep GC`.
    - GC runs automatically, identifying unreachable objects and reclaiming memory.
    - Optimized for **low latency** (important in server apps).

#### Variable Life Cycle

1. **Creation**   
    - Declared via `var`, `:=`, `new`, or `make`. Allocated on stack or heap.        
2. **Usage**    
    - Referenced by pointers, variables, or structures.        
3. **Reachability**    
    - As long as something references the variable, it stays alive.        
4. **Unreachable**    
    - When no references exist, the GC marks it as "garbage".        
5. **Collection**    
    - GC reclaims the memory.

```go
func main() {
    p := new(int)  // allocated
    *p = 42        // in use
    fmt.Println(*p)
    p = nil        // no references left
    // memory eligible for GC
}
```
#### Garbage Collection

- Go has a **garbage collector (GC)**. No manual `free()` or `delete`.
- When no references remain, memory is reclaimed automatically.
- **Tri-color algorithm** (simplified):    
    - **White set**: candidates for collection.
    - **Grey set**: objects to scan.
    - **Black set**: reachable objects.
- GC runs concurrently with program execution.

```go
func main() {
    p := new(int)
    *p = 42
    fmt.Println(*p)
    p = nil // memory eligible for GC
}
```

- You can **tune GC** with `GOGC` (garbage collection target percentage).
```bash
GOGC=100 go run main.go   # default
GOGC=off go run main.go   # disable GC (not recommended)
```

