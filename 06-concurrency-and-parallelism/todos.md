Maps and Concurrency - Maps Are NOT Thread-Safe

```go
fatal error: concurrent map read and map write
```

---
# TBD PATTERNS
## 1️⃣2️⃣ WaitGroup — Waiting for Goroutines

`var wg sync.WaitGroup  wg.Add(1) go func() {     defer wg.Done()     work() }()  wg.Wait()`

✔ Clean goroutine synchronization  
✔ Essential in production

---

## 1️⃣3️⃣ Context — Cancellation & Deadlines (VERY IMPORTANT)

### Why Context?

- Stop goroutines
    
- Propagate cancellation
    
- Handle timeouts
    

`ctx, cancel := context.WithCancel(context.Background()) defer cancel()`

---

### Context in Goroutine

`go func(ctx context.Context) {     for {         select {         case <-ctx.Done():             return         default:             work()         }     } }(ctx)`

🧠 **Interview rule**

> Every long-running goroutine should accept a context.

---

## 1️⃣4️⃣ Fan-Out / Fan-In Pattern

### Fan-Out: distribute work

`for i := 0; i < 5; i++ {     go worker(jobs, results) }`

### Fan-In: collect results

`for i := 0; i < 5; i++ {     fmt.Println(<-results) }`

Used in:

- Worker pools
    
- Parallel processing
    

---

## 1️⃣5️⃣ Worker Pool (CLASSIC INTERVIEW)

`jobs := make(chan int, 10) results := make(chan int, 10)  for i := 0; i < 3; i++ {     go worker(jobs, results) }  for j := 0; j < 5; j++ {     jobs <- j } close(jobs)  for i := 0; i < 5; i++ {     fmt.Println(<-results) }`

---

## 1️⃣6️⃣ Data Races & Safety ⚠️

### Unsafe (Race Condition)

`x := 0 go func() { x++ }() go func() { x++ }()`

🚨 Data race

---

### Safe Options

1. **Channels**
    
2. **Mutex**
    

`var mu sync.Mutex mu.Lock() x++ mu.Unlock()`

---

## 1️⃣7️⃣ Async vs Parallel (INTERVIEW QUESTION)

|Term|Meaning|
|---|---|
|Concurrency|Tasks overlap|
|Parallelism|Tasks run at same time|

Go supports **both**.

Controlled by:

`runtime.GOMAXPROCS(n)`

---

## 1️⃣8️⃣ Common Pitfalls ⚠️

1. Forgetting to wait for goroutines
    
2. Goroutine leaks
    
3. Closing channel from receiver
    
4. Writing to closed channel (panic)
    
5. Overusing buffered channels
    
6. Ignoring context cancellation
    
7. Data races
    

---

## 1️⃣9️⃣ When to Use What (Cheat Sheet)

|Use Case|Tool|
|---|---|
|Fire-and-forget|goroutine|
|Sync tasks|WaitGroup|
|Data exchange|channel|
|Cancellation|context|
|Shared state|mutex|
|Timeout|select + time.After|

---

## 🧠 Final Mental Models (MEMORIZE)

> **Goroutines are cheap**  
> **Channels synchronize + communicate**  
> **select controls async flow**  
> **context stops goroutines**  
> **Leaks are worse than panics**