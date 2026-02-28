- It is the **standard way to control goroutines lifecycle**.
- `context` provides:
	- **Cancellation**
	- **Timeouts**
	- **Deadlines**
	- **Request-scoped values**

> Every long-running goroutine must be cancelable. This is done using `context`. #InterviewQs 

## Creating Contexts

###  Root Context
```go
ctx := context.Background()
```

### Cancelable Context
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### Timeout Context
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

### Deadline Context
```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
defer cancel()
```

## Using Context in Goroutines

- Goroutine exits cleanly  
- Prevents leaks

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("stopped")
            return
        default:
            work()
        }
    }
}
```

## ctx.Done() and ctx.Err()

```go
<-ctx.Done()
```
- Blocks until canceled or timeout
- Channel-based (fits Go model)


```go
ctx.Err()
```
- Returns:
	- `context.Canceled`
	- `context.DeadlineExceeded`

## Context Propagation
#important 

- Context is **passed explicitly** down the call chain
- Context flows **top-down**, never stored globally #InterviewQs 

```go
func handler(ctx context.Context) {
    service(ctx)
}

func service(ctx context.Context) {
    dbCall(ctx)
}
```

## Context Values 
#important 

```go
ctx = context.WithValue(ctx, "userID", 42)
```

- Used for:
	- Request IDs
	- Auth metadata
	- Tracing

-  Do NOT use for:
	- Optional params
	- Business logic data

## Context + HTTP

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    process(ctx)
}
```
- Automatically canceled when:
	- Client disconnects
	- Timeout occurs

## Context vs WaitGroup
#InterviewQs 

- **They solve different problems**. Often used **together**.

| Feature             | WaitGroup | Context |
| ------------------- | --------- | ------- |
| Wait for goroutines | ✅         | ❌       |
| Cancel goroutines   | ❌         | ✅       |
| Timeout             | ❌         | ✅       |
| Propagation         | ❌         | ✅       |

## `sync` and `context` Work Together

- `context` → cancellation, `WaitGroup` → synchronization

```go
func worker(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    select {
    case <-ctx.Done():
        return
    case <-time.After(time.Second):
        fmt.Println("done")
    }
}
```

## Summary

- “Go concurrency is goroutines + channels.”
- “Mutex protects memory; context controls lifetime.”
- “Every long-running goroutine must accept a context.”
- “WaitGroup waits; context cancels.”