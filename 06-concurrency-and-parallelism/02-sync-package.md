- The `sync` package provides **low-level synchronization primitives** that allow **multiple goroutines to safely share memory**.
- Use `sync` **when goroutines need to coordinate access to shared state**.

## Sync: Use-Case

```go
x := 0

go func() { x++ }()
go func() { x++ }()
```

- Without synchronization:
	- Data Race - Two goroutines write `x` at the same time. Result is unpredictable

## Mutual Exclusion Lock - sync.Mutex

- A **mutex** ensures - Only **one goroutine** can access a critical section at a time
- Lock → exclusive access → unlock

```go
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}
```

### Benefits of Mutex

- Blocks other goroutines
- Prevents race conditions
- Very fast when uncontended

### Mistakes in Mutex

- Forgetting to unlock  
- Locking too large code blocks  
- Locking inside loops unnecessarily

- Always use `defer` when possible:
```go
mu.Lock()
defer mu.Unlock()
```

## Read/Write Lock - sync.RWMutex

- Multiple readers allowed
- Writers need exclusive access
- **Use RWMutex when**
	- Reads >> Writes
	- Heavy read-only workloads

```go
var mu sync.RWMutex
var data map[string]int

func read(key string) int {
    mu.RLock()
    defer mu.RUnlock()
    return data[key]
}

func write(key string, value int) {
    mu.Lock()
    defer mu.Unlock()
    data[key] = value
}
```

## Waiting for Goroutines - sync.WaitGroup

- Waits for all the Go routines to be finished
- Blocks until counter goes to zero  
- Essential for async workflows

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        fmt.Println(i)
    }(i)
}

wg.Wait()
```

### Working with WaitGroups

- `Add()` before goroutine starts
- Each goroutine must call `Done()`
- Do **not** reuse WaitGroup incorrectly

## Run Code Exactly Once - sync.Once

- It is used for - 
	- Singleton initialization
	- Lazy loading
	- Safe one-time setup

- Thread-safe  
- Guaranteed once, even with many goroutines

```go
var once sync.Once

func initDB() {
    fmt.Println("Initializing DB")
}

func handler() {
    once.Do(initDB)
}
```

## Concurrent Map - sync.Map

- A **lock-free concurrent map**, optimized for:
	- Many readers
	- Few writes

- Do NOT use `sync.Map` as default  
- Prefer `map + mutex`
- Use `sync.Map` only when:
	- Extremely high concurrency
	- Read-heavy workloads
	- Unknown key set

```go
var m sync.Map

m.Store("x", 10)
v, ok := m.Load("x")
m.Delete("x")
```

